package model

import (
	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/scheduler/engine"
	"github.com/Lyusis/NaotanMonitor/scheduler/queued"
	"github.com/Lyusis/NaotanMonitor/scheduler/saver"
)

// ConcurrentEngineWorker 快捷模板
var ConcurrentEngineWorker = engine.ConcurrentEngine{
	Scheduler:   &queued.Scheduler{},
	WorkerCount: conf.WorkerCount,
	SaveChan:    saver.ItemSaver(),
	RequestChan: make(chan engine.Request),
	Workers:     engine.Worker,
}

func init() {
	ConcurrentEngineWorker.Run()
}
