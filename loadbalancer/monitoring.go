package main

import (
	"fmt"
	"time"
)

type Monitor struct {
	pool []*Worker
}

func (m *Monitor) Monitor() {
	for {
		// Every 50 milliseconds monitor the workers
		time.Sleep(time.Millisecond * 50)
		load := []int{}
		for _, worker := range m.pool {
			load = append(load, worker.pending)
		}
		fmt.Printf("%v\n", load)
	}
}
