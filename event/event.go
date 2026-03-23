package event

import (
	"context"
	"time"
)

type Task struct {
	id       string
	fn       func() error
	duration time.Duration
}

type Environment struct {
	Tasks         []*Task
	stack         []*Task
	callbackQueue []*Task
	timerQueue    []*Task
}

func (e *Environment) Start() {
	if len(e.Tasks) == 0 {
		return
	}

	// other apis should send to the callback queue
	//
	callbackCh := make(chan *Task, 90)
	_ = callbackCh
}

// This is the only api that the main stack would read
// from
func (e *Environment) updateStack(ctx context.Context, cbCh <-chan *Task, outbound chan<- *Task) {
}
