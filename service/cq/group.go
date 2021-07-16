package cq

import (
	"fmt"
	"github.com/Lyusis/NaotanMonitor/utils"
	"net/http"
	"strings"

	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/logger"
)

func SendQQGroupMessage(groupId string, message string) {
	client := &http.Client{}
	urlStr := fmt.Sprintf("http://%s:%d/send_group_msg?group_id=%s&message=%s", conf.CQSendDest.IP, conf.CQSendDest.Port, groupId, message)
	if !strings.EqualFold("", conf.Token) {
		urlStr += "&access_token=" + conf.Token
	}
	logger.Sugar.Info("发送Q群消息", logger.FormatTitle("URL"), urlStr)
	fmt.Println(urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatError(requestErr))
		return
	}
	utils.BasicReceiver(client.Do(request))
}

func SendGroupMsgObserveTarget(groupId string, message string, target, from interface{}) {
	if target == from {
		SendQQGroupMessage(groupId, message)
	}
}

func SendGroupMsgObserveTargetString(groupId string, message string, target, from string) {
	if strings.Contains(target, from) {
		SendQQGroupMessage(groupId, message)
	}
}
