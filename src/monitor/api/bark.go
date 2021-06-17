package api

import (
	"monitor/config"
	"monitor/logger"
	"monitor/model"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type saberserverResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendBarkMessage(liveData model.LiveData) {

	if !config.RoomStatusList[liveData.RoomId] {
		url := "https://sctapi.ftqq.com/SCT45921Tqj6arbImzDYshqstl5siyKf9.send?title=" + config.RoomList[liveData.RoomId] + "&desp=开播啦!"
		header := "application/json;charset=UTF-8"
		postResponse, postResponseErr := http.Post(url, header, nil)
		if postResponseErr != nil {
			logger.Warn("信息发送返回json接收失败\t%+v", postResponseErr)
		}
		jsonAll, jsonErr := ioutil.ReadAll(postResponse.Body)
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