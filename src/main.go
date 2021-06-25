package main

import (
	"api"
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

	api.SendQQGroupMessage(config.GroupId, "重启中")

	fmt.Println("CQ监听服务启动中")
	go func() {
		server.NewServer(config.CQServer)
	}()

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		SaveChan:         persist.ItemSaver(),
		RequestProcessor: engine.Worker,
	}

	fmt.Println("推送服务启动中")
	for {
		utils.Delay(config.Wait)
		requestList := make([]engine.Request, 0)
		for index, name := range config.RoomList {
			url := baseurl + strconv.Itoa(index)
			requestList = append(requestList, engine.Request{
				Url:        url,
				Name:       name,
				Parser: 	live.ParseLiveData,
			})
		}
		e.Run(requestList...)
	}
}
