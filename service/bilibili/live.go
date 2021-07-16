package bilibili

import (
	"encoding/json"
	cq2 "github.com/Lyusis/NaotanMonitor/service/cq"
	saberchan2 "github.com/Lyusis/NaotanMonitor/service/saberchan"
	"net/url"

	"github.com/Lyusis/NaotanMonitor/conf"
	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/Lyusis/NaotanMonitor/scheduler/engine"
	"github.com/Lyusis/NaotanMonitor/utils"
)

func SendLiveUrl(contents []byte) engine.ResultItems {
	response := LivingUrl{}
	iveResponseErr := json.Unmarshal(contents, &response)
	if iveResponseErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Live Url } message"), logger.FormatError(iveResponseErr))
	}
	liveUrlData := response.Data
	var saveItems engine.ResultItems
	if !GetRoomStatus(liveUrlData.PlayurlInfo.Playurl.Cid) {
		return saveItems
	}

	codecs := liveUrlData.PlayurlInfo.Playurl.Stream[0].Format[0].Codec
	for _, codec := range codecs {
		if codec.CurrentQn == 10000 {
			info := codec.UrlInfo[0]
			host := info.Host
			extra := info.Extra
			baseurl := codec.BaseUrl
			urlStr := `potplayer://` + host + url.QueryEscape(baseurl+extra)
			cq2.SendQQGroupMessage(conf.GroupId, urlStr)
		}
	}
	return saveItems
}

func SendLiveStatus(contents []byte) engine.ResultItems {
	response := LiveDataResponse{}
	liveResponseErr := json.Unmarshal(contents, &response)
	if liveResponseErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Live Status } message"), logger.FormatError(liveResponseErr))
	}
	liveData := response.Data
	var saveItems engine.ResultItems

	saveItems.Items = append(saveItems.Items, liveData)
	switch liveData.LiveStatus {
	case 0, 2:
		WriteRoomStatusList(liveData.RoomId, false)
	case 1:
		if !GetRoomStatus(liveData.RoomId) {
			name := GetRoomName(liveData.RoomId)
			saberchan2.SendBarkMessage(name, "开播啦!")
			cq2.SendQQGroupMessage(conf.GroupId, utils.SingleBack(name+"开播啦!+地址:https://live.bilibili.com/", liveData.RoomId))
			//utils.Delay(2)
			SendLiveUrlService(liveData.RoomId)
		}
		WriteRoomStatusList(liveData.RoomId, true)
	}
	return saveItems
}
