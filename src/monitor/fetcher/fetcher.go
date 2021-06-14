package fetcher

import (
	"monitor/logger"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetFetcher(url string) ([]byte, error) {

	log := logger.Logger{}.InitLogger().Logger
	sugar := log.Sugar()

	response, responseError := http.Get(url)
	if responseError != nil {
		return nil, responseError
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			sugar.Warn("关闭Request失败")
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("woring status code: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
