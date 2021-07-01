package bilibili

import (
	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/monitor/engine"
	"github.com/Lyusis/NaotanMonitor/monitor/model"
	"github.com/Lyusis/NaotanMonitor/utils"
)

// SendLiveStatusService bilibili直播通知
func SendLiveStatusService() {
	const baseurl = "https://api.live.bilibili.com/room/v1/Room/room_init?id="
	go func() {
		for {
			// 获取直播状态
			requestList := make([]engine.Request, 0)
			for _, room := range RoomList {
				url := utils.SingleBack(baseurl, room.RoomId)
				requestList = append(requestList, engine.Request{
					Url:    url,
					Name:   room.Name,
					Parser: SendLiveStatus,
				})
			}
			model.ConcurrentEngineWorker.Run(requestList...)
			utils.Delay(conf.Waiting)
		}
	}()
}

// SendLiveUrlService 获取bilibili直播源
func SendLiveUrlService(roomId int) {
	const baseurl0 = "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?room_id="
	const baseurl1 = "&no_playurl=0&mask=1&qn=10000&platform=web&protocol=0,1&format=0,2&codec=0,1"
	go func() {
		// 获取直播链接
		requestList := make([]engine.Request, 0)
		requestList = append(requestList, engine.Request{
			Url:    utils.Double(baseurl0, baseurl1, roomId),
			Name:   GetRoomName(roomId),
			Parser: SendLiveUrl,
		})
		model.ConcurrentEngineWorker.Run(requestList...)
	}()
}
