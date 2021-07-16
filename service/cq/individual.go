package cq

import (
	"fmt"
	"github.com/Lyusis/NaotanMonitor/utils"
	"net/http"
	"strings"

	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/logger"
)

func SendPersonalMessage(userId string, message string) {
	client := &http.Client{}
	urlStr := fmt.Sprintf("http://%s:%d/send_private_msg?user_id=%s&message=%s", conf.CQSendDest.IP, conf.CQSendDest.Port, userId, message)
	if !strings.EqualFold("", conf.Token) {
		urlStr += "&access_token=" + conf.Token
	}
	fmt.Println(urlStr)
	logger.Sugar.Info("发送私聊消息", logger.FormatTitle("URL"), urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatError(requestErr))
		return
	}
	utils.BasicReceiver(client.Do(request))
}
