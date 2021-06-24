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
			logger.Sugar.Warn("http请求返回信息接收异常", logger.FormatTitle("Code"), resp.StatusCode)
			return
		}

		if err != nil {
			logger.Sugar.Warn("http请求返回信息接收失败", logger.FormatTitle("WRONG"), err)
			return
		}

		jsonAll, jsonErr := ioutil.ReadAll(resp.Body)
		if jsonErr != nil {
			logger.Sugar.Warn("信息发送返回json读取失败", logger.FormatTitle("WRONG"), jsonErr)
			return
		}

		jsonData := JsonData{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			logger.Sugar.Warn("信息发送返回json解析失败", logger.FormatTitle("WRONG"), unmarshalErr)
			return
		}

		if jsonData.Code != 0 {
			logger.Sugar.Warn("发送请求失败", logger.FormatTitle("WRONG"), jsonData.Message)
			return
		}
	}
	return
}
