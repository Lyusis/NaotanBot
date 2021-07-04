package cq

import (
	"github.com/Lyusis/NaotanMonitor/conf"
	"strconv"
	"strings"
)

func AJun(message MessageMessage) {
	SendGroupMsgObserveTarget(conf.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
}

func At(message MessageMessage) {
	SendGroupMsgObserveTargetString(conf.GroupId, "?", "[CQ:at,qq="+conf.QQ+"]", message.Message)
}

func AutoReturn(message MessageMessage) {
	if strings.EqualFold(message.MessageType, "private") {
		SendPersonalMessage(strconv.Itoa(message.UserId), "?")
	}
}
