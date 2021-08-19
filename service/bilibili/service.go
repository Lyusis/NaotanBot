package bilibili

import (
	"strconv"
	"strings"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/scheduler/instance"
	"github.com/Lyusis/NaotanBot/service/cq"
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

// InsertVup 添加Vup订阅
func InsertVup(msgMessage cq.MessageMessage) {
	if msgMessage.IsAt() {
		if msgMessage.IsMsgHave("订阅") {
			message := msgMessage.Message
			info := strings.Split(strings.Split(message, "订阅")[1], " ")
			nickname := info[0]
			roomId, atoiErr := strconv.Atoi(strings.TrimSpace(info[1]))
			if atoiErr != nil {
				cq.SendTool.SendGroupMessage(conf.GroupId, "订阅信息有误,房间号不应有数字以外的字符! ")
				return
			}
			sth := make([]map[string]interface{}, 0)
			liver := make(map[string]interface{}, 1)
			liver[conf.NicknameToml] = nickname
			liver[conf.RoomIdToml] = roomId
			sth = append(sth, liver)
			err := conf.AddListConfig(conf.LiversToml, sth)
			if err != nil {
				cq.SendTool.SendGroupMessage(conf.GroupId, "订阅失败, 请联系管理员! ")
			} else {
				cq.SendTool.SendGroupMessage(conf.GroupId, "订阅成功! 正在读取直播列表请稍候……")
			}
		}
	}
}

// DeleteVup 删除Vup订阅
func DeleteVup(msgMessage cq.MessageMessage) {
	if msgMessage.IsAt() {
		if msgMessage.IsMsgHave("删除") {
		}
	}
}

// SelectVup 查询Vup订阅
func SelectVup(msgMessage cq.MessageMessage) {
	if msgMessage.IsAt() {
		if msgMessage.IsMsgHave("查询") {
		}
	}
}
