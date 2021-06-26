package service

import (
	"config"
	"model"
	"monitor/engine"
	"service/bilibili"
	"utils"
)

var (
	baseurl = "https://api.live.bilibili.com/room/v1/Room/room_init?id="
)

// SendLiveStatusService bilibili直播通知
func SendLiveStatusService() {
	go func() {
		for {
			// 获取直播状态
			requestList := make([]engine.Request, 0)
			for index, name := range config.RoomList {
				url := utils.SingleBack(baseurl, index)
				requestList = append(requestList, engine.Request{
					Url:    url,
					Name:   name,
					Parser: bilibili.SendLiveStatus,
				})

			}
			model.ConcurrentEngineWorker.Run(requestList...)
			utils.Delay(config.Wait)
		}
	}()
}
