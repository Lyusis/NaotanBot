package basic

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/service/cq/server"
)

func SendUpdateMsg() {
	if "" != conf.Announcement {
		server.SendTool.SendGroupMessage(conf.GroupId, conf.Announcement)
	}
}

func SendQuitMsg() {
	server.SendTool.SendGroupMessage(conf.GroupId, conf.QuitMessage)
}
