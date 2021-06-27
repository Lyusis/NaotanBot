package model

import (
	"config"
	"monitor/engine"
	"monitor/saver"
	"monitor/scheduler"
)

// ConcurrentEngineWorker 快捷使用模板
var ConcurrentEngineWorker = engine.ConcurrentEngine{
	Scheduler:   &scheduler.QueuedScheduler{},
	WorkerCount: config.WorkerCount,
	SaveChan:    saver.ItemSaver(),
	Workers:     engine.Worker,
}
