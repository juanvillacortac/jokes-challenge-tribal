package models

import "time"

type Joke struct {
	Id           string   `json:"id"`
	Categories   []string `json:"categories"`
	CreatedAtStr string   `json:"created_at"`
	UpdatedAtStr string   `json:"updated_at"`
	IconUrl      string   `json:"icon_url"`
	Value        string   `json:"value"`
	Url          string   `json:"url"`
}

func (j *Joke) CreatedAt() time.Time {
	return time.Now()
}

func (j *Joke) UpdatedAt() time.Time {
	return time.Now()
}
