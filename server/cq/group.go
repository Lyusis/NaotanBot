package cq

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/Lyusis/NaotanMonitor/server/common"
)

func SendQQGroupMessage(groupId string, message string) {
	client := &http.Client{}
	urlStr := "http://" + conf.CQServer + ":5700/send_group_msg?group_id=" + groupId + "&message=" + message
	if !strings.EqualFold("", conf.Token) {
		urlStr += "&access_token=" + conf.Token
	}
	logger.Sugar.Info("发送Q群消息", logger.FormatTitle("URL"), urlStr)
	fmt.Println(urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatTitle("WRONG"), requestErr)
		return
	}
	common.BasicReceiver(client.Do(request))
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
