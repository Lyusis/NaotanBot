package api

import (
	"monitor/logger"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func BasicReceiver(resp *http.Response, err error) {
	if err != nil {
		logger.Warn("信息发送返回json接收失败\t%+v", err)
	}

	jsonAll, jsonErr := ioutil.ReadAll(resp.Body)
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