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

// SendGroupMsgObservePersonTarget 监控指定人是否发消息, 发送群消息
func (cqSender *Sender) SendGroupMsgObservePersonTarget(groupId string, message string, target, from interface{}) {
	if target == from {
		cqSender.SendGroupMessage(groupId, message)
	}
}

// SendGroupMsgObserveAtString 监控是否被at, 发送群消息
func (cqSender *Sender) SendGroupMsgObserveAtString(groupId string, message string, msgMessage MessageMessage) {
	if msgMessage.isAt() {
		cqSender.SendGroupMessage(groupId, message)
	}
}

// InsertVup 添加Vup订阅
func (cqSender *Sender) InsertVup(msgMessage MessageMessage) {
	if msgMessage.isAt() {
		if msgMessage.isMsgHave("订阅") {
			message := msgMessage.Message
			info := strings.Split(strings.Split(message, "订阅")[1], " ")
			nickname := info[1]
			roomId, atoiErr := strconv.Atoi(info[0])
			if atoiErr != nil {
				cqSender.SendGroupMessage(conf.GroupId, "订阅信息有误,房间号不应有数字以外的字符! ")
			}
			sth := make([]map[string]interface{}, 0)
			liver := make(map[string]interface{}, 1)
			liver[conf.NicknameToml] = nickname
			liver[conf.RoomIdToml] = roomId
			sth = append(sth, liver)
			conf.AddListConfig(conf.LiversToml, sth)
		}
	}
}

// DeleteVup 删除Vup订阅
func (cqSender *Sender) DeleteVup(msgMessage MessageMessage) {
	if msgMessage.isAt() {
		if msgMessage.isMsgHave("删除") {
		}
	}
}

// SelectVup 查询Vup订阅
func (cqSender *Sender) SelectVup(msgMessage MessageMessage) {
	if msgMessage.isAt() {
		if msgMessage.isMsgHave("查询") {
		}
	}
}

// AJun 当阿骏说话时
func (cqSender *Sender) AJun(message MessageMessage) {
	cqSender.SendGroupMsgObservePersonTarget(conf.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
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

// isMsgHave message中是否有该内容
func (msgMessage *MessageMessage) isMsgHave(shouldHave string) bool {
	if strings.Contains(msgMessage.Message, shouldHave) {
		return true
	} else {
		return false
	}
}

// isAt 判断bot是否被at
func (msgMessage *MessageMessage) isAt() bool {
	at := "[CQ:at,qq=" + conf.QQ + "]"
	return msgMessage.isMsgHave(at)
}
