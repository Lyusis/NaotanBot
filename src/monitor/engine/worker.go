package engine

import (
	"go.uber.org/zap"
	"monitor/fetcher"
	"monitor/logger"
	"monitor/utils"
)

func Worker(request Request) (Result, error) {

	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		return Result{}, bodyErr
	}

	utils.Delay()

	logger.Logger.Info("Fetching...", zap.String("Url", request.Url), zap.String("Name", request.Name))
	result := request.Parser(body)

	return result, nil
}
