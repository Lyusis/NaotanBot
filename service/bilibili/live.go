package bilibili

import (
	"encoding/json"
	"net/url"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/saberchan"
	"github.com/Lyusis/NaotanBot/utils"
)

func sendLiveUrl(contents []byte) engine.ResultItems {
	response := LivingUrl{}
	iveResponseErr := json.Unmarshal(contents, &response)
	if iveResponseErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to parsing { Live Url } message"), logger.FormatError(iveResponseErr))
	}
	liveUrlData := response.Data
	var saveItems engine.ResultItems
	if !getRoomStatus(liveUrlData.PlayurlInfo.Playurl.Cid) {
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
			cq.SendTool.SendGroupMessage(conf.GroupId, urlStr)
		}
	}
	return saveItems
}

func sendLiveStatus(contents []byte) engine.ResultItems {
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
		writeRoomStatusList(liveData.RoomId, false)
	case 1:
		if !getRoomStatus(liveData.RoomId) {
			name := getRoomName(liveData.RoomId)
			saberchan.SendBarkMessage(name, "开播啦!")
			cq.SendTool.SendGroupMessage(conf.GroupId, utils.SingleBack(name+"开播啦! 地址:https://live.bilibili.com/", liveData.RoomId))
			SendLiveUrlService(liveData.RoomId)
			writeRoomStatusList(liveData.RoomId, true)
		}
	}
	return saveItems
}
