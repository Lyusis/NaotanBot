package engine

import (
	"math/rand"
	"time"

	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/Lyusis/NaotanMonitor/monitor/fetcher"
)

func Worker(request Request) (ResultItems, error) {

	if request.Name != "" {
		logger.Sugar.Info(logger.FormatMsg("Fetching"), logger.FormatTitle("URL"), request.Url, logger.FormatTitle("Name"), request.Name)
	} else {
		logger.Sugar.Info(logger.FormatMsg("Fetching"), logger.FormatTitle("URL"), request.Url)
	}
	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		logger.Sugar.Error(logger.FormatMsg("Failed to receive request body"), bodyErr)
		logger.WriteFile(logger.FormatMsg("Writing failure information"),
			time.Now().Format(logger.TimeFormatDate)+"_fail-request-body"+string(rune(rand.Intn(19960730)))+".log", body)
		return NilResult(), bodyErr
	}

	result := request.Parser(body)

	return result, nil
}
