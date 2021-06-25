package live

import (
	"api"
	"config"
	"encoding/json"
	"logger"
	"monitor/engine"
	"monitor/model"
)


func ParseLiveData(contents []byte) engine.ResultItems {
	return SendLiveData(GetLiveData(contents))
}


func GetLiveData(contents []byte) engine.ResultItems {

	liveResponse := model.LiveResponse{}

	liveResponseErr := json.Unmarshal(contents, &liveResponse)
	if liveResponseErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Living Room } message"), logger.FormatTitle("WRONG"), liveResponseErr)
	}

	result := engine.ResultItems{}

	result.Items = append(result.Items, liveResponse.Data)

	return result
}

func SendLiveData(result engine.ResultItems) engine.ResultItems {
	var saveItems engine.ResultItems
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
