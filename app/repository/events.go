package repository

import (
	"michelfortes/concurrent-app/domain"
	"time"
)

type EventRepo interface {
	Create(event *domain.Event) (*domain.Event, error)
	GetById(id string) (*domain.Event, error)
	GetByStatus(status string) ([]*domain.Event, error)
	GetByStatusAndRegisteredBetween(status string, initTime, endTime time.Time) ([]*domain.Event, error)
	Update(event *domain.Event) (*domain.Event, error)
	Delete(id string) error
}
