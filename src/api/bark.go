package api

import (
	"config"
	"logger"
	"net/http"
)

func SendBarkMessage(title, desp string) {
	url := "https://sctapi.ftqq.com/" + config.SaberchanCode + ".send?title=" + title + "&desp=" + desp
	header := "application/json;charset=UTF-8"
	logger.Sugar.Info("发送Bark消息", logger.FormatTitle("URL"), url)
	BasicReceiver(http.Post(url, header, nil))
}
