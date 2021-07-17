package common

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

var MsgBoxChan = make(chan MsgBox, conf.WorkerCount)
