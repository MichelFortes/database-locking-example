package domain

import "time"

type Event struct {
	Id        string    `json:"id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Attempts  int       `json:"attempts"`
	Status    string    `json:"status"`
}
