package main

import (
	"michelfortes/concurrent-app/config"
	"michelfortes/concurrent-app/db"
	"michelfortes/concurrent-app/worker"
	"sync"
)

const workers = 3

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

	wg := sync.WaitGroup{}

	wg.Add(workers)

	for i := 1; i <= workers; i++ {
		go worker.New(i, pool, &wg)
	}

	wg.Wait()
}
