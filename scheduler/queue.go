package scheduler

import "web-crawler/engine"
// All workers own their in chan, they have one out chan
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request // channel of worker in channels
}
// Submit job to this engine
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r // Just like simple engine
}

// worker call this function to add them to the ready queue
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			// Get the top of two queue, put them into select section
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}



