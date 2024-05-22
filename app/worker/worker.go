package worker

import (
	"context"
	"fmt"
	"michelfortes/concurrent-app/domain"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(workerId int, pool *pgxpool.Pool) {

	rows, e := pool.Query(context.Background(), "select id, payload, registered_at, attempts, status from events where status=$1 for update skip locked limit 3", "idle")
	if e != nil {
		panic(e)
	}

	events, e := pgx.AppendRows[domain.Event](make([]domain.Event, len(rows.RawValues())), rows, func(row pgx.CollectableRow) (domain.Event, error) {

		var id, payload, status string
		var registeredAt time.Time
		var attempts int

		err := row.Scan(&id, &payload, &registeredAt, &attempts, &status)

		event := domain.Event{
			Id:           id,
			Payload:      payload,
			Status:       status,
			RegisteredAt: registeredAt,
			Attempts:     attempts,
		}

		return event, err
	})
	if e != nil {
		panic(e)
	}

	fmt.Printf("\n===== results from worker %d =====\n", workerId)
	for _, e := range events {
		fmt.Println(e.Payload)
	}
}
