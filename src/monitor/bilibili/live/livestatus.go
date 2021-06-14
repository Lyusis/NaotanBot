package live

import (
	"monitor/engine"
	"monitor/logger"
	"monitor/model"
	"encoding/json"
)

func GetLiveData(contents []byte) engine.Result {
	log := logger.Logger{}.InitLogger().Logger
	sugar := log.Sugar()

	liveResponse := model.LiveResponse{}

	liveResponseErr := json.Unmarshal(contents, &liveResponse)
	if liveResponseErr != nil {
		sugar.Warn(liveResponseErr)
	}

	result := engine.Result{}

	result.Items = append(result.Items, liveResponse.Data)

	return result
}
