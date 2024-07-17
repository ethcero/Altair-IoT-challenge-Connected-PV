package scheduler

import (
	"time"
)

type Scheduler struct {
	interval int // in seconds
	ticker   *time.Ticker
}

func NewScheduler(interval int) *Scheduler {
	return &Scheduler{
		interval: interval,
		ticker:   time.NewTicker(time.Duration(interval) * time.Second),
	}
}

func (s *Scheduler) Start(task func()) {
	println("Starting scheduler")

	go func() {
		for {
			select {
			case <-s.ticker.C:
				go task()
				break
			}

		}
	}()
}

func (s *Scheduler) Stop() {
	println("Stopping scheduler")
	s.ticker.Stop()
}
