package cq

import (
	"config"
	"fmt"
	"logger"
	"net/http"
	"server/common"
	"strconv"
	"strings"
)

func SendPersonalMessage(userId int, message string) {
	client := &http.Client{}
	urlStr := "http://" + config.CQServer + ":5700/send_private_msg?user_id=" + strconv.Itoa(userId) + "&message=" + message
	if !strings.EqualFold("", config.Token) {
		urlStr += "&access_token=" + config.Token
	}
	fmt.Println(urlStr)
	logger.Sugar.Info("发送私聊消息", logger.FormatTitle("URL"), urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatTitle("WRONG"), requestErr)
		return
	}
	common.BasicReceiver(client.Do(request))
}
