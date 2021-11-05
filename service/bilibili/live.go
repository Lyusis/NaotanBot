package bilibili

import (
	"encoding/json"
	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/utils"
)

func sendLiveStatus(contents []byte) engine.ResultItems {
	var (
		saveItems engine.ResultItems
		response  LiveDataResponse
		err       error
	)

	err = json.Unmarshal(contents, &response)
	if err != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Live Status } message"), logger.FormatError(err))
	} else {
		liveData := response.Data
		saveItems.Items = append(saveItems.Items, liveData)
		for _, liveItem := range liveData {
			switch liveItem.LiveStatus {
			case 0, 2:
				writeRoomStatusList(liveItem.Uid, false)
			case 1:
				if !getRoomStatus(liveItem.Uid) {
					name := getRoomName(liveItem.Uid)
					if name == "" {
						name = liveItem.Uname
					}
					cq.SendTool.SendGroupMessage(conf.GroupId, utils.SingleBackInt(name+"开播啦!\n"+liveItem.Title+"\n"+"地址:https://live.bilibili.com/", liveItem.RoomId))
					//SendLiveUrlService(liveData.Uid)
					writeRoomStatusList(liveItem.Uid, true)
				}
			}
		}
	}

	return saveItems
}
