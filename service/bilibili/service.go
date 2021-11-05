package bilibili

import (
	"container/list"
	"encoding/json"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/scheduler/fetcher"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/scheduler/instance"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/redis"
	"github.com/Lyusis/NaotanBot/utils"
)

// SendLiveStatusService bilibili直播通知
func SendLiveStatusService() {
	const baseurl = "https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids"
	go func() {
		for {
			// 获取直播状态
			if Reload {
				Reload = false
				reloadRoomList()
			}
			<-utils.Delay(conf.Waiting)
			if len(LiverList) == 0 {
				continue
			}
			body := make(map[string][]string)
			for _, room := range LiverList {
				body["uids"] = append(body["uids"], strconv.Itoa(room.Liver.Uid))
			}
			bytesData, err := json.Marshal(body)
			if err != nil {
				logger.Sugar.Warn(logger.FormatMsg("Failed to marshal json"), logger.FormatError(err))
				return
			}
			instance.ConcurrentEngineWorker.RequestChan <- engine.Request{
				Url:    baseurl,
				Method: http.MethodPost,
				Name:   "获取主播信息",
				Body:   bytesData,
				Parser: sendLiveStatus,
			}
		}
	}()
}

// InsertVup 添加Vup订阅
func InsertVup(msgMessage cq.MessageMessage) {
	commands := "订阅 {uid} {nickname}"
	msgMessage.AtFilter(commands, func(params *list.List, SendTool cq.Sender) (string, error) {
		var (
			err      error
			nickname string
			uid      int
		)
		uid, err = strconv.Atoi(strings.TrimSpace(utils.PopUp(params)))
		if err != nil {
			return "订阅信息有误,uid不应有数字以外的字符! ", err
		}
		nickname = utils.PopUp(params)
		if nickname == "" {
			url := "https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids"
			body := make(map[string][]string)
			body["uids"] = append(body["uids"], strconv.Itoa(uid))
			var (
				response  LiveDataResponse
				bytesData []byte
				contents  []byte
			)
			bytesData, err = json.Marshal(body)
			if err == nil {
				contents, err = fetcher.Fetcher(url, http.MethodPost, bytesData)
				if err == nil {
					err = json.Unmarshal(contents, &response)
					if err == nil {
						if name := response.Data[strconv.Itoa(uid)].Uname; name != "" {
							err = redis.SetAdd("Liver", utils.SingleFrontInt(uid, ":"+name))
						}
					}
				}
			}
		} else {
			err = redis.SetAdd("Liver", utils.SingleFrontInt(uid, ":"+nickname))
		}
		Reload = true
		return "订阅失败, 请联系管理员!", err
	})
}

// DeleteVup 删除Vup订阅
func DeleteVup(msgMessage cq.MessageMessage) {
	commands := "删除订阅 {keyword}"
	msgMessage.AtFilter(commands, func(params *list.List, SendTool cq.Sender) (string, error) {
		var (
			member string
		)
		member = utils.PopUp(params)
		Reload = true
		return "删除失败, 请联系管理员! ",
			redis.SetDelete("Liver", member)
	})
}

// SelectVup 查询Vup订阅
func SelectVup(msgMessage cq.MessageMessage) {
	commands := "订阅列表"
	msgMessage.SingleAtFilter(commands, func(SendTool cq.Sender) {
		var (
			livers []string
			result string
			err    error
		)
		livers, err = redis.SetGet("Liver")
		if err != nil {
			cq.SendTool.SendGroupMessage(conf.GroupId, "无法获取订阅列表")
			return
		}
		for _, item := range livers {
			result += item + "\n"
		}
		cq.SendTool.SendGroupMessage(conf.GroupId, result)
	})
}
