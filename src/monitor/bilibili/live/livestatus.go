package live

import (
	"encoding/json"
	"monitor/engine"
	"monitor/logger"
	"monitor/model"
)

func GetLiveData(contents []byte) engine.Result {

	liveResponse := model.LiveResponse{}

	liveResponseErr := json.Unmarshal(contents, &liveResponse)
	if liveResponseErr != nil {
		logger.Sugar.Warn("直播信息json解析错误", logger.FormatTitle("WRONG"), liveResponseErr)
	}

	result := engine.Result{}

	result.Items = append(result.Items, liveResponse.Data)

	return result
}
