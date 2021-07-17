package httpserver

import (
	"encoding/json"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/cq/model"
	"io/ioutil"
	"net/http"
)

func Handler(_ http.ResponseWriter, r *http.Request) {
	eventMessage := cq.MetaEventMessage{}
	message := cq.MessageMessage{}
	postType := cq.PostType{}
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
		model.SendTool.AJun(message)
		model.SendTool.At(message)
		model.SendTool.AutoReturn(message)
	}
}
