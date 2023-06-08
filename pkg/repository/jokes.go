package repository

import (
	"challenge/pkg/models"
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/goccy/go-json"
	"golang.org/x/sync/errgroup"
)

type JokeRepository struct {
	Url string
}

func NewJokeRepository(url string) *JokeRepository {
	return &JokeRepository{
		Url: url,
	}
}

func (repo *JokeRepository) List(ctx context.Context) ([]models.Joke, error) {
	jokes := make([]models.Joke, 0)

	err := repo.GetJokes(ctx, &jokes, 25, 25)

	return jokes, err
}

func (repo *JokeRepository) GetJokes(ctx context.Context, jokes *[]models.Joke, initial, n int) error {
	if len(*jokes) >= n {
		return nil
	}

	temp := make([]models.Joke, len(*jokes))
	copy(temp, *jokes)

	var m sync.Mutex
	var eg errgroup.Group

	for i := 0; i < n; i++ {
		eg.Go(func() error {
			req, err := http.NewRequest("GET", repo.Url, nil)
			if err != nil {
				return err
			}
			req = req.WithContext(ctx)

			client := http.DefaultClient
			res, err := client.Do(req)
			if err != nil {
				return err
			}

			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)

			var joke models.Joke
			if err := json.Unmarshal(body, &joke); err != nil {
				return err
			}
			m.Lock()
			temp = append(temp, joke)
			m.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	allKeys := make(map[string]bool)
	list := []models.Joke{}
	for _, item := range temp {
		if _, value := allKeys[item.Id]; !value {
			allKeys[item.Id] = true
			list = append(list, item)
		}
	}

	if len(list) < initial {
		if err := repo.GetJokes(ctx, &list, initial, initial-len(list)); err != nil {
			return nil
		}
	}

	*jokes = list

	return nil
}
