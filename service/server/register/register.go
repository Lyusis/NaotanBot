package register

import (
	"github.com/Lyusis/NaotanBot/service/cq"

	"github.com/Lyusis/NaotanBot/service/basic"
	biliServer "github.com/Lyusis/NaotanBot/service/bilibili/server"
	dailiesServer "github.com/Lyusis/NaotanBot/service/dailylife/server"
	friendsServer "github.com/Lyusis/NaotanBot/service/friends/server"
)

func Register(message cq.MessageMessage) {
	basic.Menu(message)
	biliServer.InsertVup(message)
	biliServer.SelectVup(message)
	biliServer.DeleteVup(message)
	dailiesServer.GetWeather(message)
	dailiesServer.GetNews(message)
	friendsServer.AJun(message)
	friendsServer.InitiativeAJun(message)
	friendsServer.InsertAJW(message)
	friendsServer.Tiangou(message)
}
