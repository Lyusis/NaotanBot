package bilibili

import (
	"api/cq"
	"api/saberchan"
	"config"
	"encoding/json"
	"logger"
	"model"
	"monitor/engine"
	"utils"
)

var (
	baseurl21 = "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo?room_id="
	baseurl22 = "&no_playurl=0&mask=1&qn=10000&platform=web&protocol=0,1&format=0,2&codec=0,1"
)

func SendLiveUrl(contents []byte) engine.ResultItems {
	response := model.LivingUrl{}
	iveResponseErr := json.Unmarshal(contents, &response)
	if iveResponseErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Live Url } message"), logger.FormatTitle("WRONG"), iveResponseErr)
	}
	liveUrlData := response.Data
	var saveItems engine.ResultItems
	if !config.RoomStatusList[liveUrlData.PlayurlInfo.Playurl.Cid] {
		return saveItems
	}

	codecs := liveUrlData.PlayurlInfo.Playurl.Stream[0].Format[0].Codec
	for _, codec := range codecs {
		if codec.CurrentQn == 10000 {
			info := codec.UrlInfo[0]
			host := info.Host
			extra := info.Extra
			baseurl := codec.BaseUrl
			url := `potplayer://` + host + baseurl + extra
			cq.SendQQGroupMessage(config.GroupId, url)
		}
	}

	return saveItems
}

func SendLiveUrlService(roomId int) {
	go func() {
		// 获取直播链接
		requestList := make([]engine.Request, 0)
		requestList = append(requestList, engine.Request{
			Url:    utils.Double(baseurl21, baseurl22, roomId),
			Name:   config.RoomList[roomId],
			Parser: SendLiveUrl,
		})
		model.ConcurrentEngineWorker.Run(requestList...)
	}()
}

func SendLiveStatus(contents []byte) engine.ResultItems {
	response := model.LiveDataResponse{}
	liveResponseErr := json.Unmarshal(contents, &response)
	if liveResponseErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Live Status } message"), logger.FormatTitle("WRONG"), liveResponseErr)
	}
	liveData := response.Data
	var saveItems engine.ResultItems

	saveItems.Items = append(saveItems.Items, liveData)
	switch liveData.LiveStatus {
	case 0:
		setRoomStatusFalse(liveData.RoomId)
	case 1:
		if !config.RoomStatusList[liveData.RoomId] {
			name := config.RoomList[liveData.RoomId]
			saberchan.SendBarkMessage(name, "开播啦!")
			cq.SendQQGroupMessage(config.GroupId, utils.SingleBack(name+"开播啦!\nhttps://live.bilibili.com/", liveData.RoomId))
			SendLiveUrlService(liveData.RoomId)
		}
		setRoomStatusTrue(liveData.RoomId)
	case 2:
		setRoomStatusFalse(liveData.RoomId)
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
