package db

import (
	"context"
	"fmt"
	"michelfortes/concurrent-app/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnPoll(conf config.Config) (*pgxpool.Pool, error) {
	cpConfig, e := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", conf.User, conf.Pass, conf.DB))
	if e != nil {
		panic(e)
	}
	return pgxpool.NewWithConfig(context.Background(), cpConfig)
}
