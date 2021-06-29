package saberchan

import (
	"fmt"
	"net/http"

	"github.com/Lyusis/NaotanMonitor/config"
	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/Lyusis/NaotanMonitor/server/common"
)

func SendBarkMessage(title, desp string) {
	urlStr := "https://sctapi.ftqq.com/" + config.SaberchanCode + ".send?title=" + title + "&desp=" + desp
	fmt.Println(urlStr)
	header := "application/json;charset=UTF-8"
	logger.Sugar.Info(logger.FormatMsg("Sending BARK messages"), logger.FormatTitle("URL"), urlStr)
	common.BasicReceiver(http.Post(urlStr, header, nil))
}
