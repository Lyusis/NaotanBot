package cq

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/utils"
)

var (
	at = "[CQ:at,qq=" + strconv.Itoa(conf.QQ) + "]"
)

type Sender struct {
	SendMessage SendMessage
}

type SendMessage interface {
	SendGroupMessage(groupId int, message string)
	SendPrivateMessage(userId int, message string)
}

func (cqSender *Sender) SendGroupMessage(groupId int, message string) {
	cqSender.SendMessage.SendGroupMessage(groupId, message)
}

func (cqSender *Sender) SendPrivateMessage(groupId int, message string) {
	cqSender.SendMessage.SendPrivateMessage(groupId, message)
}

// SendGroupMsgObservePersonTarget 监控指定人是否发消息, 发送群消息
func (cqSender *Sender) SendGroupMsgObservePersonTarget(groupId int, message string, target, from interface{}) {
	if target == from {
		cqSender.SendGroupMessage(groupId, message)
	}
}

// SendGroupMsgObserveAtString 监控是否被at, 发送群消息
func (cqSender *Sender) SendGroupMsgObserveAtString(groupId int, message string, msgMessage MessageMessage) {
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
		cqSender.SendPrivateMessage(message.UserId, "?")
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
	return msgMessage.IsMsgHave(at)
}

// AtFilter at过滤器
func (msgMessage *MessageMessage) AtFilter(commands string,
	todo func(params *list.List, sender Sender) (string, error)) {
	// AT检测
	if msgMessage.IsAt() {
		// 获取数据
		var (
			message     = msgMessage.Message
			commandList = strings.Split(commands, " ")
			inputList   = utils.ExtractContent(message)

			params list.List
		)
		// 数据处理及具体实施
		inputList = inputList[1:]
		if len(commandList) < len(inputList) {
			return
		}
		for index, input := range inputList {
			command := commandList[index]
			if utils.CheckCurlyBraces(command) {
				params.PushBack(input)
			} else {
				if input == commandList[index] {
					continue
				} else {
					return
				}
			}
		}

		resultMsg, err := todo(&params, SendTool)
		if err != nil {
			SendTool.SendGroupMessage(conf.GroupId, resultMsg)
		}
		SendTool.SendGroupMessage(conf.GroupId, "操作完成! ")
	}

}

func (msgMessage *MessageMessage) SingleAtFilter(commands string, todo func(sender Sender)) {
	commandList := utils.ExtractContent(commands)
	// AT检测
	if msgMessage.IsAt() {
		// 获取数据
		var (
			message   = msgMessage.Message
			inputList = utils.ExtractContent(message)
		)
		// 数据处理及具体实施
		inputList = inputList[1:]
		if len(commandList) != len(inputList) {
			return
		}
		if inputList[0] == commandList[0] {
			todo(SendTool)
		}
	}
}
