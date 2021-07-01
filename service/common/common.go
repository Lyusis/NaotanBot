package common

import (
	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/server/cq"
)

func SendUpdateMsg() {
	cq.SendQQGroupMessage(conf.GroupId, conf.Announcement)
}
