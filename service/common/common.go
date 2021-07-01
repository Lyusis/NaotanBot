package common

import (
	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/sender/cq"
	"net/url"
)

func SendUpdateMsg() {
	cq.SendQQGroupMessage(conf.GroupId, url.QueryEscape(conf.Announcement))
}
