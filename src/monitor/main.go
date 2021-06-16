package main

import (
	"monitor/bilibili/live"
	"monitor/config"
	"monitor/engine"
	"monitor/persist"
	"monitor/scheduler"
	"monitor/utils"
	"strconv"
)

var (
	baseurl = "https://api.live.bilibili.com/room/v1/Room/room_init?id="
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         persist.ItemSaver(),
		RequestProcessor: engine.Worker,
	}

	for {
		utils.Delay()
		for index, name := range config.RoomList {
			url := baseurl + strconv.Itoa(index)
			e.Run(engine.Request{
				Url:    url,
				Name:   name,
				Parser: live.GetLiveData,
			})
		}
	}
}
