package engine

import (
	"math/rand"
	"monitor/fetcher"
	"monitor/logger"
	"time"
)

func Worker(request Request) (Result, error) {

	if request.Name != "" {
		logger.Info("Fetching", true, "URL", request.Url, "Name", request.Name)
	} else {
		logger.Info("Fetching", true, "URL", request.Url)
	}
	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		logger.Error("获取请求体失败", false, bodyErr)
		logger.WriteFile("写入失败的请求体",
			time.Now().Format(logger.TimeFormatDate)+"_fail-request-body"+string(rune(rand.Intn(19960730)))+".log", body)
		return NilResult(), bodyErr
	}

	result := request.Parser(body)

	return result, nil
}
