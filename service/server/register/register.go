package register

import (
	"github.com/Lyusis/NaotanBot/service/basic"
	biliServer "github.com/Lyusis/NaotanBot/service/bilibili/server"
	"github.com/Lyusis/NaotanBot/service/cq"
)

func Register(message cq.MessageMessage) {
	basic.AJun(message)
	biliServer.InsertVup(message)
}
