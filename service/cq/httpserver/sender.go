package httpserver

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/utils"
)

type HttpSender struct{}

func (sender *HttpSender) SendGroupMessage(groupId string, message string) {
	client := &http.Client{}
	urlStr := fmt.Sprintf("http://%s:%d/send_group_msg?group_id=%s&message=%s", conf.CQSendDest.IP, conf.CQSendDest.Port, groupId, message)
	if !strings.EqualFold("", conf.Token) {
		urlStr += "&access_token=" + conf.Token
	}
	logger.Sugar.Info("发送Q群消息", logger.FormatTitle("URL"), urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatError(requestErr))
		return
	}
	utils.BasicReceiver(client.Do(request))
}

func (sender *HttpSender) SendPrivateMessage(userId string, message string) {
	client := &http.Client{}
	urlStr := fmt.Sprintf("http://%s:%d/send_private_msg?user_id=%s&message=%s", conf.CQSendDest.IP, conf.CQSendDest.Port, userId, message)
	if !strings.EqualFold("", conf.Token) {
		urlStr += "&access_token=" + conf.Token
	}
	logger.Sugar.Info("发送私聊消息", logger.FormatTitle("URL"), urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatError(requestErr))
		return
	}
	utils.BasicReceiver(client.Do(request))
}
