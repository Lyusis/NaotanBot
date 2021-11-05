package server

import (
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/dailylife"
)

func Clock() {
	dailylife.Clock()
}

func GetWeather(msgMessage cq.MessageMessage) {
	dailylife.GetWeather(msgMessage)
}

func GetNews(msgMessage cq.MessageMessage) {
	dailylife.GetNews(msgMessage)
}
