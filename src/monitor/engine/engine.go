package engine

import (
	"fmt"
	"monitor/api"
	"monitor/config"
	"monitor/model"
)

// Run 多线程调度引擎/**
func (engine *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan Result)
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkerCount; i++ {
		engine.createWorker(engine.Scheduler.WorkerChan(), out, engine.Scheduler)
	}

	for _, request := range seeds {
		engine.Scheduler.Submit(request)
	}

	result := <-out
	for _, item := range result.Items {
		go func(item interface{}) {
			if liveData, ok := item.(model.LiveData); ok {
				fmt.Print(config.RoomList[liveData.RoomId] + ": ")
				switch liveData.LiveStatus {
				case 0:
					fmt.Println("尚未直播")
					setRoomStatusFalse(liveData)
				case 1:
					fmt.Println("直播中")
					api.SendBarkMessage(liveData)
					api.SendQQMessage(liveData)
					setRoomStatusTrue(liveData)
				case 2:
					fmt.Println("轮播中")
					setRoomStatusFalse(liveData)
				}
			}
			// TODO: 将数据放入数据库
			engine.ItemChan <- item

		}(item)
	}
}

func setRoomStatusFalse(liveData model.LiveData) {
	if config.RoomStatusList[liveData.RoomId] {
		config.RoomStatusList[liveData.RoomId] = false
	}
}

func setRoomStatusTrue(liveData model.LiveData) {
	if !config.RoomStatusList[liveData.RoomId] {
		config.RoomStatusList[liveData.RoomId] = true
	}
}