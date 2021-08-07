package saberchan

import (
	"github.com/Lyusis/NaotanBot/utils"
	"net/http"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
)

func SendBarkMessage(title, desp string) {
	if "" != conf.SaberchanCode {
		urlStr := "https://sctapi.ftqq.com/" + conf.SaberchanCode + ".send?title=" + title + "&desp=" + desp
		header := "application/json;charset=UTF-8"
		logger.Sugar.Info(logger.FormatMsg("Sending BARK messages"), logger.FormatTitle("URL"), urlStr)
		utils.BasicReceiver(http.Post(urlStr, header, nil))
	}
}
