package fetcher

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
)

func Fetcher(url string, method string, body []byte) ([]byte, error) {

	var (
		ctx     = context.Background()
		request *http.Request
		err     error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(conf.Waiting)*time.Second)
	defer cancel()

	if body != nil {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	} else {
		request, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		logger.Sugar.Warn(logger.FormatMsg("No responded request"), logger.FormatError(err))
		return nil, err
	}
	request = request.WithContext(ctx)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed response"), logger.FormatError(err))
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received HTTP request exception, Code: %d", response.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Sugar.Warn(logger.FormatMsg("Failed to close request"), logger.FormatError(err))
		}
	}(response.Body)

	return ioutil.ReadAll(response.Body)
}
