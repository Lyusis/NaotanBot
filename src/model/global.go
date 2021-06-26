package model

import (
	"config"
	"monitor/engine"
	"monitor/persist"
	"monitor/scheduler"
)

var ConcurrentEngineWorker = engine.ConcurrentEngine{
	Scheduler:   &scheduler.QueuedScheduler{},
	WorkerCount: config.WorkerCount,
	SaveChan:    persist.ItemSaver(),
	Workers:     engine.Worker,
}
