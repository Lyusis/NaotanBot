package engine

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/utils"
)

// Run 调度引擎/**
func (engine *ConcurrentEngine) Run() {

	out := make(chan ResultItems, conf.WorkerCount)
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkerCount; i++ {
		engine.createWorker(out, engine.Scheduler)
	}

	for {
		select {
		case request := <-engine.RequestChan:
			go func(request Request) {
				engine.Scheduler.Submit(request)
				result := <-out
				// TODO: 将数据放入数据库
				for _, item := range result.Items {
					if item != DelayOp {
						engine.SaveChan <- item
					} else {
						// Magic Count
						conf.Waiting = utils.LongDelay
					}
				}
			}(request)
		}
	}
}

// createWorker 创建Worker/**
func (engine *ConcurrentEngine) createWorker(
	out chan ResultItems, ready Scheduler) {
	go func() {
		for {
			in := ready.WorkerChan()
			ready.WorkerReady(in)
			request := <-in
			go engine.Workers(request, out)
		}
	}()
}
