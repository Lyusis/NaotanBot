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
	SendGroupMessage(groupId string, message string)
	SendPrivateMessage(userId string, message string)
}

func (cqSender *Sender) SendGroupMessage(groupId string, message string) {
	cqSender.SendMessage.SendGroupMessage(groupId, message)
}

func (cqSender *Sender) SendPrivateMessage(groupId string, message string) {
	cqSender.SendMessage.SendPrivateMessage(groupId, message)
}

// SendGroupMsgObserveTarget 与target一致
func (cqSender *Sender) SendGroupMsgObserveTarget(groupId string, message string, target, from interface{}) {
	if target == from {
		cqSender.SendGroupMessage(groupId, message)
	}
}

// SendGroupMsgObserveTargetString 存在target
func (cqSender *Sender) SendGroupMsgObserveTargetString(groupId string, message string, target, from string) {
	if strings.Contains(target, from) {
		cqSender.SendGroupMessage(groupId, message)
	}
}

func (cqSender *Sender) AJun(message MessageMessage) {
	cqSender.SendGroupMsgObserveTarget(conf.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
}

func (cqSender *Sender) At(message MessageMessage) {
	cqSender.SendGroupMsgObserveTargetString(conf.GroupId, "?", "[CQ:at,qq="+conf.QQ+"]", message.Message)
}

func (cqSender *Sender) AutoReturn(message MessageMessage) {
	if strings.EqualFold(message.MessageType, "private") {
		cqSender.SendPrivateMessage(strconv.Itoa(message.UserId), "?")
	}
}
