package api

import (
	"encoding/json"
	"io/ioutil"
	"monitor/logger"
	"net/http"
)

func BasicReceiver(resp *http.Response, err error) {
	if resp != nil {
		if err != nil {
			logger.Warn("信息发送返回json接收失败", false, err)
		}

		jsonAll, jsonErr := ioutil.ReadAll(resp.Body)
		if jsonErr != nil {
			logger.Warn("信息发送返回json读取失败", false, jsonErr)
		}

		jsonData := saberserverResponse{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			logger.Warn("信息发送返回json解析失败", false, unmarshalErr)
		}

		if jsonData.Code != 0 {
			logger.Warn("发送请求失败", false, jsonData.Message)
		}
	}
}
