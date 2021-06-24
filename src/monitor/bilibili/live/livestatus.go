package live

import (
	"api"
	"config"
	"encoding/json"
	"logger"
	"monitor/engine"
	"monitor/model"
)

func GetLiveData(contents []byte) engine.ResultItems {

	liveResponse := model.LiveResponse{}

	liveResponseErr := json.Unmarshal(contents, &liveResponse)
	if liveResponseErr != nil {
		logger.Sugar.Warn("直播信息json解析错误", logger.FormatTitle("WRONG"), liveResponseErr)
	}

	result := engine.ResultItems{}

	result.Items = append(result.Items, liveResponse.Data)

	return result
}

func SendLiveData(result engine.ResultItems) engine.SaveItems {
	var saveItems engine.SaveItems
	for i := 0; i < len(result.Items); i++ {
		if liveData, ok := result.Items[i].(model.LiveData); ok {
			saveItems.Items = append(saveItems.Items, liveData)
			name := config.RoomList[liveData.RoomId]

			switch liveData.LiveStatus {
			case 0:
				setRoomStatusFalse(liveData.RoomId)
			case 1:
				if !config.RoomStatusList[liveData.RoomId] {
					api.SendBarkMessage(name, "开播啦!")
					api.SendQQGroupMessage(config.GroupId, name+"开播啦!")
				}
				setRoomStatusTrue(liveData.RoomId)
			case 2:
				setRoomStatusFalse(liveData.RoomId)
			}
		}
	}
	return saveItems
}

func setRoomStatusFalse(roomId int) {
	if config.RoomStatusList[roomId] {
		config.RoomStatusList[roomId] = false
	}
}

func setRoomStatusTrue(roomId int) {
	if !config.RoomStatusList[roomId] {
		config.RoomStatusList[roomId] = true
	}
}
