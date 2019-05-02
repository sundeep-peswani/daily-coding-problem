package main

import (
	"fmt"
	"time"
)

// A Scheduler to run timer events
type Scheduler struct {
	inQueue int

	queued    chan int
	completed chan int
}

func main() {
	s := newScheduler()
	exit := make(chan int)

	go s.run(exit)

	s.schedule(func() { fmt.Println("Called second, after 2s") }, 2000)
	s.schedule(func() { fmt.Println("Called first, after 500ms") }, 500)

	<-exit
}

func newScheduler() *Scheduler {
	var s Scheduler
	s.queued = make(chan int)
	s.completed = make(chan int)

	return &s
}

func (s *Scheduler) schedule(f func(), n int) *time.Timer {
	s.queued <- 0

	return time.AfterFunc(time.Duration(n)*time.Millisecond, func() {
		f()
		s.completed <- 0
	})
}

func (s *Scheduler) run(exit chan int) {
	for {
		select {
		case <-s.queued:
			s.inQueue++
		case <-s.completed:
			s.inQueue--
			if s.inQueue == 0 {
				exit <- 0
				return
			}
		}
	}
}
