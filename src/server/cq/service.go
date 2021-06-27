package cq

import (
	"config"
)

func AJun(message MessageMessage) {
	SendQQGroupMsgObserveTarget(config.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
}

func At(message MessageMessage) {
	SendQQGroupMsgObserveTargetString(config.GroupId, "?", "[CQ:at,qq="+config.QQ+"]", message.Message)
}
