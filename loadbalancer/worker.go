package main

type Worker struct {
	requests chan Request // Channel containing requests buffered channel
	pending  int          // number of pending tasks
	index    int          // index in the heap
	id       int
}

func (w *Worker) work(done chan *Worker) {
	for {
		// Consume a request off the channel
		req := <-w.requests
		// Execute the operation and place back in the request
		req.Ch <- req.Fn()
		// Tell the channel of workers that this one is done with a task
		done <- w
	}
}
