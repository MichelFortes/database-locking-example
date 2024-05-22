package main

import (
	"michelfortes/concurrent-app/config"
	"michelfortes/concurrent-app/db"
	"michelfortes/concurrent-app/worker"
	"time"
)

func main() {

	cfg, e := config.FromArgs()
	if e != nil {
		panic(e)
	}

	pool, e := db.NewConnPoll(cfg)
	if e != nil {
		panic(e)
	}
	defer pool.Close()

	for i := 1; i <= 3; i++ {
		go worker.New(i, pool)
	}

	time.Sleep(5 * time.Second)
}
