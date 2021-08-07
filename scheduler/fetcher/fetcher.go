package fetcher

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
)

func GetFetcher(url string) ([]byte, error) {

	ctx := context.Background()
	// TODO: 超时时间可配置化
	ctx, cancel := context.WithTimeout(ctx, time.Duration(conf.Waiting)*time.Second)
	defer cancel()

	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("No responded request"), logger.FormatError(requestErr))
		return nil, requestErr
	}

	request = request.WithContext(ctx)

	response, responseErr := http.DefaultClient.Do(request)
	if responseErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed response"), logger.FormatError(responseErr))
		return nil, responseErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to close request"), logger.FormatError(err))
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received HTTP request exception, Code: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
