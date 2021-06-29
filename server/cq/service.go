package cq

import (
	"github.com/Lyusis/NaotanMonitor/config"
)

func AJun(message MessageMessage) {
	SendGroupMsgObserveTarget(config.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
}

func At(message MessageMessage) {
	SendGroupMsgObserveTargetString(config.GroupId, "?", "[CQ:at,qq="+config.QQ+"]", message.Message)
}

func AutoReturn(message MessageMessage) {
	SendPersonalMessage(message.UserId, "?")
}
