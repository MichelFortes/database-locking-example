package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"michelfortes/concurrent-app/ini"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Event struct {
	Id        string    `json:"id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Attempts  int       `json:"attempts"`
	Status    string    `json:"status"`
}

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

	rows, e := pool.Query(context.Background(), "select id, payload, created_at, updated_at, attempts, status from events where status=$1", "idle")
	if e != nil {
		panic(e)
	}

	events, e := pgx.AppendRows[Event](make([]Event, len(rows.RawValues())), rows, func(row pgx.CollectableRow) (Event, error) {

		var id, payload, status string
		var created, updated time.Time
		var attempts int

		err := row.Scan(&id, &payload, &created, &updated, &attempts, &status)

		event := Event{
			Id:        id,
			Payload:   payload,
			Status:    status,
			CreatedAt: created,
			UpdatedAt: updated,
			Attempts:  attempts,
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
