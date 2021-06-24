package main

import (
	"config"
	"fmt"
	"monitor/bilibili/live"
	"monitor/engine"
	"monitor/persist"
	"monitor/scheduler"
	"server"
	"strconv"
	"utils"
)

var (
	baseurl = "https://api.live.bilibili.com/room/v1/Room/room_init?id="
)

func main() {

	fmt.Println("CQ监听服务启动中")	
	go func() { 
		server.NewServer(config.CQServer) 
	}()

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         persist.ItemSaver(),
		RequestProcessor: engine.Worker,
	}

	fmt.Println("推送服务启动中")	
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
