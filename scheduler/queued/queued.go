package queued

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
)

type Scheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *Scheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *Scheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *Scheduler) WorkerReady(
	w chan engine.Request) {
	s.workerChan <- w
}

func (s *Scheduler) Run() {
	s.workerChan = make(chan chan engine.Request, conf.WorkerCount)
	s.requestChan = make(chan engine.Request, conf.WorkerCount)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 &&
				len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
