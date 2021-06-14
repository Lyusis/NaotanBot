package engine

import (
	"monitor/config"
	"monitor/logger"
	"monitor/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Run 多线程调度引擎/**
func (engine *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan Result)
	engine.Scheduler.Run()
	roomStatusList := make(map[int]bool)

	for i := 0; i < engine.WorkerCount; i++ {
		engine.createWorker(engine.Scheduler.WorkerChan(), out, engine.Scheduler)
	}

	for _, request := range seeds {
		engine.Scheduler.Submit(request)
	}

	for index := range config.RoomList {
		roomStatusList[index] = false
	}

	result := <-out
	for _, item := range result.Items {
		go func(item interface{}) {
			if liveData, ok := item.(model.LiveData); ok {
				fmt.Print(config.RoomList[liveData.RoomId] + ": ")
				switch liveData.LiveStatus {
				case 0:
					fmt.Println("尚未直播")
					setRoomStatusFalse(roomStatusList, liveData)
				case 1:
					sendMessage(liveData, roomStatusList)
					setRoomStatusTrue(roomStatusList, liveData)
				case 2:
					fmt.Println("轮播中")
					setRoomStatusFalse(roomStatusList, liveData)
				}
			}
			engine.ItemChan <- item
			// TODO: 将数据放入数据库

		}(item)
	}

}

func setRoomStatusFalse(roomStatusList map[int]bool, liveData model.LiveData) {
	if roomStatusList[liveData.RoomId] {
		roomStatusList[liveData.RoomId] = false
	}
}

func setRoomStatusTrue(roomStatusList map[int]bool, liveData model.LiveData) {
	if !roomStatusList[liveData.RoomId] {
		roomStatusList[liveData.RoomId] = true
	}
}

func sendMessage(liveData model.LiveData, roomStatusList map[int]bool) {
	log := logger.Logger{}.InitLogger().Logger
	sugar := log.Sugar()

	fmt.Println("直播中")
	if !roomStatusList[liveData.RoomId] {
		url := "https://sctapi.ftqq.com/SCT45921Tqj6arbImzDYshqstl5siyKf9.send?title=" + config.RoomList[liveData.RoomId] + "&desp=开播啦!"
		header := "application/json;charset=UTF-8"
		postResponse, postResponseErr := http.Post(url, header, nil)
		if postResponseErr != nil {
			sugar.Warn(postResponseErr)
		}
		jsonAll, jsonErr := ioutil.ReadAll(postResponse.Body)
		if jsonErr != nil {
			sugar.Warn(jsonErr)
		}

		jsonData := saberserverResponse{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			sugar.Warn(unmarshalErr)
			log.Warn("json解析失败")
		}

		if jsonData.Code != 0 {
			sugar.Warnf("发送请求失败, 失败理由: %s", jsonData.Message)
		}
	}
}
