package main

import (
	"context"
	"log"
	"time"
)

func (loop *EventLoop) runTimerQueue(ctx context.Context, wait <-chan *Task) <-chan *Task {
	callback_ch := make(chan *Task)
	go func() {
		defer close(callback_ch)
		for {
			select {
			case <-ctx.Done():
				return
			case t, open := <-wait:
				if !open {
					return
				}
				if t.duration == nil {
					log.Fatalf("received task.duration as nil: %+v\n", t)
					return
				}
				timer := time.NewTimer(*t.duration)
				<-timer.C
				// would make more sense to add this to an actual 'array' or 'stack' out <- t
				t.duration = nil
				callback_ch <- t
			}
		}
	}()
	return callback_ch
}

func (loop EventLoop) execute() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	waitCh := make(chan *Task)
	toStackCh := loop.runTimerQueue(ctx, waitCh)

	mainCh := make(chan *Task)
	go loop.pushToStack2(ctx, toStackCh, mainCh)
	defer close(waitCh)

	executedTasks := 0
	totalTasks := len(loop.Tasks)
	for _, task := range loop.Tasks {
		if task.duration != nil {
			waitCh <- task
			continue
		}

		task.fn()
		executedTasks++
	}
	for {
		if executedTasks == totalTasks {
			break
		}

		readyTask := <-mainCh
		readyTask.fn()
		executedTasks++
	}

}
func (loop EventLoop) pushToStack2(ctx context.Context, others <-chan *Task, main chan<- *Task) {
	for {
		select {
		case <-ctx.Done():
			return
		case t, open := <-others:
			if !open {
				return
			}
			main <- t
		}
	}
}
