package main

import (
	"config"
	"monitor/bilibili/live"
	"monitor/engine"
	"monitor/persist"
	"monitor/scheduler"
	"strconv"
	"utils"
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
		utils.Delay(0)
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
