package common

import (
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/service/cq"
	"net/url"
)

func SendUpdateMsg() {
	cq.SendQQGroupMessage(conf.GroupId, url.QueryEscape(conf.Announcement))
}
