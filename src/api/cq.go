package api

import (
	"config"
	"encoding/json"
	"io/ioutil"
	"monitor/logger"
	"net/http"
)

func CQHandler(r *http.Request) {
	eventMessage := MetaEventMessage{}
	message := MessageMessage{}
	postType := PostType{}
	readAll, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Sugar.Warn("信息发送返回json读取失败", logger.FormatTitle("WRONG"), err)
	}
	checkErr := json.Unmarshal(readAll, &postType)
	if checkErr != nil {
		logger.Sugar.Warn("信息发送返回json解析失败", logger.FormatTitle("WRONG"), checkErr)
	}

	switch postType.PostType {
	case "meta_event":
		jsonErr := json.Unmarshal(readAll, &eventMessage)
		if jsonErr != nil {
			logger.Sugar.Warn("cq请求上报返回json解析失败", logger.FormatTitle("WRONG"), jsonErr)
		}
		//fmt.Printf("%+v\n", message)
	case "message":
		jsonErr := json.Unmarshal(readAll, &message)
		if jsonErr != nil {
			logger.Sugar.Warn("cq消息上报返回json解析失败", logger.FormatTitle("WRONG"), jsonErr)
		}
		SendQQGroupMsgObserveTarget(config.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
	}
}

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
	BasicReceiver(client.Do(request))
}

func SendQQGroupMsgObserveTarget(groupId string, message string, target, from int) {
	if target == from {
		SendQQGroupMessage(groupId, message)
	}
}
