package api

import (
	"config"
	"monitor/logger"
	"net/http"
)

func SendBarkMessage(title, desp string) {
	url := "https://sctapi.ftqq.com/" + config.SaberchanCode + ".send?title=" + title + "&desp=" + desp
	header := "application/json;charset=UTF-8"
	logger.Info("发送Bark消息", true, "URL", url)
	BasicReceiver(http.Post(url, header, nil))
}
