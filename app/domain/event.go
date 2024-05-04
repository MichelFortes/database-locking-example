package domain

import "time"

type Event struct {
	Id           string    `json:"id"`
	Payload      string    `json:"payload"`
	RegisteredAt time.Time `json:"registered_at"`
	Attempts     int       `json:"attempts"`
	Status       string    `json:"status"`
}
