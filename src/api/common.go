package api

import (
	"encoding/json"
	"io/ioutil"
	"monitor/logger"
	"net/http"
)

type JsonData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BasicReceiver(resp *http.Response, err error) {
	if resp != nil {
		if resp.StatusCode != 200 {
			logger.Warn("http请求返回信息接收异常", true, "Code", resp.StatusCode)
			return
		}

		if err != nil {
			logger.Warn("http请求返回信息接收失败", false, err)
			return
		}

		jsonAll, jsonErr := ioutil.ReadAll(resp.Body)
		if jsonErr != nil {
			logger.Warn("信息发送返回json读取失败", false, jsonErr)
			return
		}

		jsonData := JsonData{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			logger.Warn("信息发送返回json解析失败", false, unmarshalErr)
			return
		}

		if jsonData.Code != 0 {
			logger.Warn("发送请求失败", false, jsonData.Message)
			return
		}
	}
	return
}
