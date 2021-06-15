package fetcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"monitor/logger"
	"net/http"
)

func GetFetcher(url string) ([]byte, error) {

	response, responseError := http.Get(url)
	if responseError != nil {
		return nil, responseError
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Logger.Sugar().Warn("关闭Request失败")
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("woring status code: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
