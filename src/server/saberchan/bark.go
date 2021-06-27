package saberchan

import (
	"config"
	"fmt"
	"logger"
	"net/http"
	"server/common"
)

func SendBarkMessage(title, desp string) {
	urlStr := "https://sctapi.ftqq.com/" + config.SaberchanCode + ".send?title=" + title + "&desp=" + desp
	fmt.Println(urlStr)
	header := "application/json;charset=UTF-8"
	logger.Sugar.Info(logger.FormatMsg("Sending BARK messages"), logger.FormatTitle("URL"), urlStr)
	common.BasicReceiver(http.Post(urlStr, header, nil))
}
