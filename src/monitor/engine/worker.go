package engine

import (
	"monitor/fetcher"
)

func Worker(request Request) (Result, error) {

	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		return Result{}, bodyErr
	}

	result := request.Parser(body)

	return result, nil
}
