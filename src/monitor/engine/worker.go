package engine

import (
	"math/rand"
	"monitor/fetcher"
	"monitor/logger"
	"time"
)

func Worker(request Request) (Result, error) {

	if request.Name != "" {
		logger.Sugar.Info("Fetching", logger.FormatTitle("URL"), request.Url, logger.FormatTitle("Name"), request.Name)
	} else {
		logger.Sugar.Info("Fetching", logger.FormatTitle("URL"), request.Url)
	}
	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		logger.Sugar.Error("获取请求体失败", bodyErr)
		logger.WriteFile("写入失败的请求体",
			time.Now().Format(logger.TimeFormatDate)+"_fail-request-body"+string(rune(rand.Intn(19960730)))+".log", body)
		return NilResult(), bodyErr
	}

	result := request.Parser(body)

	return result, nil
}
