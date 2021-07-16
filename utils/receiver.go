package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Lyusis/NaotanBot/logger"
)

type JsonData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BasicReceiver(resp *http.Response, err error) {
	if resp != nil {
		if err != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to receive HTTP request"), logger.FormatError(err))
			return
		}

		if resp.StatusCode != http.StatusOK {
			logger.Sugar.Warn(logger.FormatMsg("Received HTTP request exception"), logger.FormatTitle("Code"), resp.StatusCode)
			return
		}

		jsonAll, jsonErr := ioutil.ReadAll(resp.Body)
		defer func(Body io.ReadCloser) {
			if bodyCloseErr := Body.Close(); bodyCloseErr != nil {
				logger.Sugar.Panic(logger.FormatMsg("Failed to close response body"), bodyCloseErr)
			}
		}(resp.Body)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to read JSON message"), logger.FormatError(jsonErr))
			return
		}

		jsonData := JsonData{}
		unmarshalErr := json.Unmarshal(jsonAll, &jsonData)
		if unmarshalErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to parse JSON"), logger.FormatError(unmarshalErr))
			return
		}

		if jsonData.Code != 0 {
			logger.Sugar.Warn(logger.FormatMsg("Failed to send request"), logger.FormatTitle("WRONG"), jsonData.Message)
			return
		}
	}
	return
}
