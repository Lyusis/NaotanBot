package model

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/scheduler/queued"
	"github.com/Lyusis/NaotanBot/scheduler/saver"
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
	go func() {
		ConcurrentEngineWorker.Run()
	}()
}
