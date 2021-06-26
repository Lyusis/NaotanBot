package engine

import (
	"config"
	"logger"
)

// Run 调度引擎/**
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
			// TODO: 将数据放入数据库
			for _, item := range result.Items {
				engine.SaveChan <- item
			}
		}(request)
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
			result, err := engine.Workers(request)
			if err != nil {
				logger.Sugar.Warn(logger.FormatMsg("Failed to create Worker"), err)
				return
			}
			out <- result
		}
	}()
}
