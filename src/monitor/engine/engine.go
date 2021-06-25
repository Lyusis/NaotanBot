package engine

import (
	"config"
	"logger"
)

// Run 多线程调度引擎/**
func (engine *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ResultItems, config.WorkerCount)
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkerCount; i++ {
		engine.createWorker(out, engine.Scheduler)
	}

	for _, request := range seeds {
		go func(request Request) {
			engine.Scheduler.Submit(request)
			result := <-out
			items := request.PostParser(result)
			// TODO: 将数据放入数据库
			for _, item := range items.Items {
				engine.SaveChan <- item
			}
		}(request)
	}
}

// createWorker Worker创建/**
func (engine *ConcurrentEngine) createWorker(
	out chan ResultItems, ready Scheduler) {
	go func() {

		// FIXME: 协程未关闭
		in := ready.WorkerChan()
		ready.WorkerReady(in)
		request := <-in
		result, err := engine.RequestProcessor(request)
		if err != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to create Worker"), err)
			return
		}

		out <- result
	}()
}
