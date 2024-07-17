package scheduler

import (
	"testing"
	"time"
)

func TestNewScheduler(t *testing.T) {
	s := NewScheduler(5)
	if s.interval != 5 {
		t.Errorf("Expected interval to be 5, got %d", s.interval)
	}
}

func TestScheduler_Start(t *testing.T) {
	s := NewScheduler(1)
	taskCalled := false
	task := func() {
		taskCalled = true
	}
	s.Start(task)
	time.Sleep(2 * time.Second)
	if !taskCalled {
		t.Errorf("Expected task to be called")
	}
	s.Stop()
}
