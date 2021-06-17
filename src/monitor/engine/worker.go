package engine

import (
	"monitor/fetcher"
	"monitor/logger"
)

func Worker(request Request) (Result, error) {

	logger.Info("Fetching...\t| Url: %s\t| Name: %s", request.Url, request.Name)
	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		return Result{}, bodyErr
	}

	result := request.Parser(body)

	return result, nil
}
