package bilibili

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/scheduler/instance"
	"github.com/Lyusis/NaotanBot/utils"
)

// SendLiveStatusService bilibili直播通知
func SendLiveStatusService() {
	const baseurl = "https://api.live.bilibili.com/room/v1/Room/room_init?id="
	go func() {
		for {
			// 获取直播状态
			if conf.ReLoad {
				conf.CheckedReload()
				reloadRoomList()
			}
			<-utils.Delay(conf.Waiting)
			for _, room := range RoomList {
				url := utils.SingleBack(baseurl, room.Liver.RoomId)
				instance.ConcurrentEngineWorker.RequestChan <- engine.Request{
					Url:    url,
					Name:   room.Liver.Nickname,
					Parser: sendLiveStatus,
				}
			}
		}
	}()
}

// SendLiveUrlService 获取bilibili直播源
func SendLiveUrlService(roomId int) {
	const baseurl0 = "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?room_id="
	const baseurl1 = "&no_playurl=0&mask=1&qn=10000&platform=web&protocol=0,1&format=0,2&codec=0,1"
	go func() {
		// 获取直播链接
		instance.ConcurrentEngineWorker.RequestChan <- engine.Request{
			Url:    utils.Double(baseurl0, baseurl1, roomId),
			Name:   getRoomName(roomId),
			Parser: sendLiveUrl,
		}
	}()
}
