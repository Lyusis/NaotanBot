package api

import (
	"config"
	"encoding/json"
	"io/ioutil"
	"logger"
	"net/http"
)

func CQHandler(r *http.Request) {
	eventMessage := MetaEventMessage{}
	message := MessageMessage{}
	postType := PostType{}
	readAll, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Sugar.Warn("Server failed to read JSON message", logger.FormatTitle("WRONG"), err)
	}
	checkErr := json.Unmarshal(readAll, &postType)
	if checkErr != nil {
		logger.Sugar.Warn("Server failed to parse JSON message(TYPE)", logger.FormatTitle("WRONG"), checkErr)
	}

	switch postType.PostType {
	case "meta_event":
		jsonErr := json.Unmarshal(readAll, &eventMessage)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(META_EVENT)"), logger.FormatTitle("WRONG"), jsonErr)
		}
		//fmt.Printf("%+v\n", message)
	case "message":
		// TODO: 剥离
		jsonErr := json.Unmarshal(readAll, &message)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(MESSAGE)"), logger.FormatTitle("WRONG"), jsonErr)
		}
		SendQQGroupMsgObserveTarget(config.GroupId, "阿骏不要再舔了", 1565255741, message.UserId)
	}
}



func SendQQGroupMessage(groupId string, message string) {
	client := &http.Client{}
	url := "http://" + config.CQServer + ":5700/send_group_msg?group_id=" + groupId + "&message=" + message + "&access_token=" + "guanrenchi"
	logger.Sugar.Info(logger.FormatMsg("Sending QQ-Group message"), logger.FormatTitle("URL"), url)
	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to send QQ-Group message failed"), logger.FormatTitle("WRONG"), requestErr)
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
