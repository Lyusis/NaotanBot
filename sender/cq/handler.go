package cq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Lyusis/NaotanMonitor/logger"
)

func Handler(r *http.Request) {
	eventMessage := MetaEventMessage{}
	message := MessageMessage{}
	postType := PostType{}
	readAll, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Sugar.Warn("Server failed to read JSON message", logger.FormatError(err))
	}
	checkErr := json.Unmarshal(readAll, &postType)
	if checkErr != nil {
		logger.Sugar.Warn("Server failed to parse JSON message(TYPE)", logger.FormatError(checkErr))
	}

	switch postType.PostType {
	case "meta_event":
		jsonErr := json.Unmarshal(readAll, &eventMessage)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(META_EVENT)"), logger.FormatError(jsonErr))
		}
	case "message":
		jsonErr := json.Unmarshal(readAll, &message)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(MESSAGE)"), logger.FormatError(jsonErr))
		}
		AJun(message)
		At(message)
		AutoReturn(message)
	}
}
