package server

import (
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/friends"
)

func InsertAJW(msgMessage cq.MessageMessage) {
	friends.InsertAJW(msgMessage)
}
func AJun(msgMessage cq.MessageMessage) {
	friends.AJun(msgMessage)
}
func InitiativeAJun(msgMessage cq.MessageMessage) {
	friends.InitiativeAJun(msgMessage)
}

func Tiangou(msgMessage cq.MessageMessage) {
	friends.Tiangou(msgMessage)
}
