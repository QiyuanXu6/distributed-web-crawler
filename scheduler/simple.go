package scheduler

import "web-crawler/engine"
// All workers share the same in chan and out chan
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (scheduler *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		scheduler.workerChan <- r
	}()
}

func (scheduler *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	scheduler.workerChan = in
}