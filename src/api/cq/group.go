package cq

import (
	"api"
	"config"
	"logger"
	"net/http"
)

func SendQQGroupMessage(groupId string, message string) {
	client := &http.Client{}
	url := "http://" + config.CQServer + ":5700/send_group_msg?group_id=" + groupId + "&message=" + message + "&access_token=" + "guanrenchi"
	logger.Sugar.Info("发送Q群消息", logger.FormatTitle("URL"), url)
	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatTitle("WRONG"), requestErr)
		return
	}
	//request.Header.Add("AccessToken", "guanrenchi")
	api.BasicReceiver(client.Do(request))
}

func SendQQGroupMsgObserveTarget(groupId string, message string, target, from int) {
	if target == from {
		SendQQGroupMessage(groupId, message)
	}
}
