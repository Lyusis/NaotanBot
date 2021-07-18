package basic

import (
	"github.com/Lyusis/NaotanBot/conf"
	server2 "github.com/Lyusis/NaotanBot/service/cq/server"
)

func SendUpdateMsg() {
	server2.SendTool.SendGroupMessage(conf.GroupId, conf.Announcement)
}
