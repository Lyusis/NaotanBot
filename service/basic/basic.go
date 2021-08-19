package basic

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/service/cq"
)

func SendUpdateMsg() {
	if "" != conf.Announcement {
		cq.SendTool.SendGroupMessage(conf.GroupId, conf.Announcement)
	}
}

func SendQuitMsg() {
	cq.SendTool.SendGroupMessage(conf.GroupId, conf.QuitMessage)
}

func SendDelayMsg() {
	if "" != conf.Announcement {
		cq.SendTool.SendGroupMessage(conf.GroupId, "风控, 暂停服务30分钟")
	}
}

func AJun(message cq.MessageMessage) {
	if "" != conf.Announcement {
		cq.SendTool.SendGroupMsgObservePersonTarget(conf.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
	}
}
