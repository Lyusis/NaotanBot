package common

import (
	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/service/cq"
	"net/url"
)

func SendUpdateMsg() {
	cq.SendQQGroupMessage(conf.GroupId, url.QueryEscape(conf.Announcement))
}
