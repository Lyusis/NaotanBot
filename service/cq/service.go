package cq

import (
	"strconv"
	"strings"

	"github.com/Lyusis/NaotanBot/conf"
)

type Sender struct {
	SendMessage SendMessage
}

type SendMessage interface {
	SendGroupMessage(groupId, message string)
	SendPrivateMessage(userId, message string)
}

func (cqSender *Sender) SendGroupMessage(groupId, message string) {
	cqSender.SendMessage.SendGroupMessage(groupId, message)
}

func (cqSender *Sender) SendPrivateMessage(groupId, message string) {
	cqSender.SendMessage.SendPrivateMessage(groupId, message)
}

// SendGroupMsgObservePersonTarget 监控指定人是否发消息, 发送群消息
func (cqSender *Sender) SendGroupMsgObservePersonTarget(groupId, message string, target, from interface{}) {
	if target == from {
		cqSender.SendGroupMessage(groupId, message)
	}
}

// SendGroupMsgObserveAtString 监控是否被at, 发送群消息
func (cqSender *Sender) SendGroupMsgObserveAtString(groupId, message string, msgMessage MessageMessage) {
	if msgMessage.IsAt() {
		cqSender.SendGroupMessage(groupId, message)
	}
}

// At 当Bot被at时
func (cqSender *Sender) At(message MessageMessage) {
	cqSender.SendGroupMsgObserveAtString(conf.GroupId, "?", message)
}

// AutoReturn 当Bot被私聊时
func (cqSender *Sender) AutoReturn(message MessageMessage) {
	if strings.EqualFold(message.MessageType, "private") {
		cqSender.SendPrivateMessage(strconv.Itoa(message.UserId), "?")
	}
}

// IsMsgHave message中是否有该内容
func (msgMessage *MessageMessage) IsMsgHave(shouldHave string) bool {
	if strings.Contains(msgMessage.Message, shouldHave) {
		return true
	} else {
		return false
	}
}

// IsAt 判断bot是否被at
func (msgMessage *MessageMessage) IsAt() bool {
	at := "[CQ:at,qq=" + conf.QQ + "]"
	return msgMessage.IsMsgHave(at)
}
