package cq

import (
	"github.com/Lyusis/NaotanBot/conf"
)

const (
	SendGroup   = "send_group_msg"
	SendPrivate = "send_private_msg"
)

type MsgBox struct {
	Action  string
	Id      string
	Message string
}

var (
	MsgBoxChan = make(chan MsgBox, conf.WorkerCount)
	SendTool   = Sender{}
)

func SetHttpSendTool() {
	SendTool = Sender{
		SendMessage: &HttpSender{},
	}
}

func SetWSSendTool() {
	SendTool = Sender{
		SendMessage: &WSSender{},
	}
}
