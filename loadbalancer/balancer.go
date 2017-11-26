package main

import (
	"container/heap"
)

type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker // Channel of workers that have completed their request
	work chan Request // Channel of requests
}

func (b *Balancer) Balance(work chan Request) {
	for {
		select {
		case req := <-b.work: // Received some request
			b.dispatch(req) // Send to a worker
		case w := <-b.done: // Worker finished
			b.complete(w) // Mark as completed
		}
	}
}

func (b *Balancer) dispatch(req Request) {
	// Grab lightest load worker
	w := heap.Pop(&b.pool).(*Worker)
	// Send over the work
	w.requests <- req
	// Increment pending
	w.pending++
	// Put it back into the heap
	heap.Push(&b.pool, w)
}

func (b *Balancer) complete(w *Worker) {
	w.pending--                   // One less pending request
	heap.Remove(&b.pool, w.index) // Remove it from the heap
	heap.Push(&b.pool, w)         // Add it back the heap
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	worker := x.(*Worker)
	worker.index = n
	*p = append(*p, worker)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	worker := old[n-1]
	worker.index = -1 // for safety
	*p = old[0 : n-1]
	return worker
}
