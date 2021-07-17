package model

import (
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/cq/wsserver"
)

var SendTool = cq.Sender{
	SendMessage: &wsserver.WSSender{},
}
