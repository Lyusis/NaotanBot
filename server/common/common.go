package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Lyusis/NaotanMonitor/logger"
)

type JsonData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BasicReceiver(resp *http.Response, err error) {
	if resp != nil {
		if err != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to receive HTTP request"), logger.FormatTitle("WRONG"), err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			logger.Sugar.Warn(logger.FormatMsg("Received HTTP request exception"), logger.FormatTitle("Code"), resp.StatusCode)
			return
		}

		jsonAll, jsonErr := ioutil.ReadAll(resp.Body)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to read JSON message"), logger.FormatTitle("WRONG"), jsonErr)
			return
		}

		jsonData := JsonData{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to parse JSON"), logger.FormatTitle("WRONG"), unmarshalErr)
			return
		}

		if jsonData.Code != 0 {
			logger.Sugar.Warn(logger.FormatMsg("Failed to send request"), logger.FormatTitle("WRONG"), jsonData.Message)
			return
		}
	}
	return
}
