package main

import (
	"math/rand"
	"time"
)

type Request struct {
	Fn func() int // Function this request should perform
	Ch chan int   // Channel the result should be sent on
}

func requester(id int, work chan<- Request) {
	// Each requester has a channel where they can see the work that's been done
	c := make(chan int)
	for {
		// Every 50 milliseconds send in some work
		time.Sleep(time.Millisecond * 50)

		// Send in the Request to the work channel
		work <- Request{
			// Add two random numbers
			Fn: func() int { return rand.Intn(10) + rand.Intn(10) },
			Ch: c,
		}
		_ = <-c
	}
}
