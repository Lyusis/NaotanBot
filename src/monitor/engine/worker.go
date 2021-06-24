package engine

import (
	"monitor/fetcher"
	"monitor/logger"
)

func Worker(request Request) (Result, error) {

	logger.Sugar.Info("Fetching", logger.FormatTitle("URL"), request.Url, logger.FormatTitle("Name"), request.Name)
	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		return Result{}, bodyErr
	}

	result := request.Parser(body)

	return result, nil
}
