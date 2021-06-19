package api

import (
	"config"
	"net/http"
)

func SendQQGroupMessage(groupId string, message string) {
	url := "https://" + config.CQServer + ":5700/send_group_msg?group_id=" + groupId + "&message=" + message
	BasicReceiver(http.Get(url))
}
