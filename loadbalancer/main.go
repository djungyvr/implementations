package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Channel to put finished workers
	done := make(chan *Worker)

	// Create 4 workers
	workers := []*Worker{}

	for i := 0; i < 4; i++ {
		worker := &Worker{
			// Each worker gets its own channel
			requests: make(chan Request, 1000),
			id:       i,
		}
		workers = append(workers, worker)
		go worker.work(done) // Each worker should also publish their results to the done channel
	}

	// requesters will publish to the work channel
	// load balancer will subscribe from work channel
	work := make(chan Request)

	// Initialize the load balancer with 4 workers
	b := Balancer{
		pool: workers,
		done: done,
		work: work,
	}

	// Monitors the load on the workers
	m := Monitor{
		pool: workers,
	}

	// b will consume from the work channel and send off to workers
	go b.Balance(work)

	//
	go m.Monitor()

	// create 100 requesters requests
	for i := 0; i < 100; i++ {
		go requester(i, work)
	}

	fmt.Println("Press enter to exit")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
