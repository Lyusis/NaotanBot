package server

import (
	"github.com/Lyusis/NaotanBot/service/bilibili"
	"github.com/Lyusis/NaotanBot/service/cq"
)

func SendLiveStatusService() {
	bilibili.SendLiveStatusService()
}

func InsertVup(msgMessage cq.MessageMessage) {
	bilibili.InsertVup(msgMessage)
}

func SelectVup(msgMessage cq.MessageMessage) {
	bilibili.SelectVup(msgMessage)
}

func DeleteVup(msgMessage cq.MessageMessage) {
	bilibili.DeleteVup(msgMessage)
}
