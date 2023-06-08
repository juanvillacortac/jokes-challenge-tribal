package handlers

import (
	"challenge/pkg/repository"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetJokes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	url := os.Getenv("DATA_URL")

	repo := repository.NewJokeRepository(url)

	log.Println(fmt.Sprintf("Obteniendo bromas desde %s", url))

	data, err := repo.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(data)
}
