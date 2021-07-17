package wsserver

import (
	"github.com/Lyusis/NaotanBot/service/common"
)

type WSSender struct{}

func (sender *WSSender) SendGroupMessage(groupId string, message string) {
	msgBox := common.MsgBox{
		Action: common.SendGroup, Id: groupId, Message: message,
	}
	common.MsgBoxChan <- msgBox
}

func (sender *WSSender) SendPrivateMessage(userId string, message string) {
	msgBox := common.MsgBox{
		Action: common.SendPrivate, Id: userId, Message: message,
	}
	common.MsgBoxChan <- msgBox
}
