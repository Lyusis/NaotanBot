package api

import (
	"monitor/config"
	"monitor/logger"
	"monitor/model"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func SendQQMessage(liveData model.LiveData) {
	if !config.RoomStatusList[liveData.RoomId] {
		url := "http://127.0.0.1:5700/send_group_msg?group_id=540419281&message=" + config.RoomList[liveData.RoomId] + "开播啦"
		getResponse, getResponseErr := http.Get(url)
		if getResponseErr != nil {
			logger.Warn("信息发送返回json接收失败\t%+v", getResponseErr)
		}

		jsonAll, jsonErr := ioutil.ReadAll(getResponse.Body)
		if jsonErr != nil {
			logger.Warn("信息发送返回json读取失败\t%+v", jsonErr)
		}

		jsonData := saberserverResponse{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			logger.Warn("信息发送返回json解析失败\t%+v", unmarshalErr)
		}

		if jsonData.Code != 0 {
			logger.Warn("发送请求失败\t%+v", jsonData.Message)
		}

	}
}