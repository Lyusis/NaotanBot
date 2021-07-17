package basic

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/server"
)

func SendUpdateMsg() {
	server.SendTool.SendGroupMessage(conf.GroupId, conf.Announcement)
}
