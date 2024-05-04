package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"michelfortes/concurrent-app/domain"
	"michelfortes/concurrent-app/ini"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var confPath string
	flag.StringVar(&confPath, "c", "", "-c [config file]")
	flag.Parse()

	iniFile, e := ini.From(confPath)
	if e != nil {
		panic(e)
	}

	cpConfig, e := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", iniFile.User, iniFile.Pass, iniFile.DB))
	if e != nil {
		panic(e)
	}

	pool, e := pgxpool.NewWithConfig(context.Background(), cpConfig)
	if e != nil {
		panic(e)
	}
	defer pool.Close()

	rows, e := pool.Query(context.Background(), "select id, payload, registered_at, attempts, status from events where status=$1", "idle")
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

	b, e := json.Marshal(events)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))
}
