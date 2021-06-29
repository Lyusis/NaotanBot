package model

import (
	"github.com/Lyusis/NaotanMonitor/config"
	"github.com/Lyusis/NaotanMonitor/monitor/engine"
	"github.com/Lyusis/NaotanMonitor/monitor/saver"
	"github.com/Lyusis/NaotanMonitor/monitor/scheduler"
)

// ConcurrentEngineWorker 快捷使用模板
var ConcurrentEngineWorker = engine.ConcurrentEngine{
	Scheduler:   &scheduler.QueuedScheduler{},
	WorkerCount: config.WorkerCount,
	SaveChan:    saver.ItemSaver(),
	Workers:     engine.Worker,
}
