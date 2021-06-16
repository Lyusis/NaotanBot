package engine

import (
	"go.uber.org/zap"
	"monitor/fetcher"
	"monitor/logger"
)

func Worker(request Request) (Result, error) {

	logger.Logger.Info("Fetching...", zap.String("Url", request.Url), zap.String("Name", request.Name))
	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		return Result{}, bodyErr
	}

	result := request.Parser(body)

	return result, nil
}
