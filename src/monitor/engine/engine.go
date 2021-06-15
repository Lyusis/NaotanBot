package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"monitor/config"
	"monitor/logger"
	"monitor/model"
	"net/http"
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
					sendMessage(liveData)
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

func sendMessage(liveData model.LiveData) {

	fmt.Println("直播中")
	if !config.RoomStatusList[liveData.RoomId] {
		url := "https://sctapi.ftqq.com/SCT45921Tqj6arbImzDYshqstl5siyKf9.send?title=" + config.RoomList[liveData.RoomId] + "&desp=开播啦!"
		header := "application/json;charset=UTF-8"
		postResponse, postResponseErr := http.Post(url, header, nil)
		if postResponseErr != nil {
			logger.Logger.Sugar().Warn(postResponseErr)
		}
		jsonAll, jsonErr := ioutil.ReadAll(postResponse.Body)
		if jsonErr != nil {
			logger.Logger.Sugar().Warn(jsonErr)
		}

		jsonData := saberserverResponse{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			logger.Logger.Sugar().Warn(unmarshalErr)
			logger.Logger.Warn("json解析失败")
		}

		if jsonData.Code != 0 {
			logger.Logger.Sugar().Warnf("发送请求失败, 失败理由: %s", jsonData.Message)
		}
	}
}
