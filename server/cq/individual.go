package cq

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/Lyusis/NaotanMonitor/server/common"
)

func SendPersonalMessage(userId int, message string) {
	client := &http.Client{}
	urlStr := "http://" + conf.CQSendDest.IP + ":" + conf.CQSendDest.Port + "/send_private_msg?user_id=" + strconv.Itoa(userId) + "&message=" + message
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
	common.BasicReceiver(client.Do(request))
}
