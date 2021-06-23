package fetcher

import (
	"config"
	"fmt"
	"io"
	"io/ioutil"
	"monitor/logger"
	"net/http"
	"time"
)

func GetFetcher(url string) ([]byte, error) {

	client := http.Client{
		Timeout: time.Duration(config.Wait/5) * time.Second,
	}

	response, responseError := client.Get(url)
	if responseError != nil {
		logger.Error("未响应请求或响应失败, 将在下次轮询中再次请求", false, responseError)
		return nil, responseError
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Warn("关闭Request失败", false, err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败, 错误代码: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
