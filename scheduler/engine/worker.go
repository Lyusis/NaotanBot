package engine

import (
	"strings"

	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/scheduler/fetcher"
)

func Worker(request Request, resultItemsChan chan ResultItems) {

	if request.Name != "" {
		logger.Sugar.Info(logger.FormatMsg("Fetching"), logger.FormatTitle("URL"), request.Url, logger.FormatTitle("NickName"), request.Name)
	} else {
		logger.Sugar.Info(logger.FormatMsg("Fetching"), logger.FormatTitle("URL"), request.Url)
	}
	body, bodyErr := fetcher.GetFetcher(request.Url)
	if bodyErr != nil {
		logger.Sugar.Error(logger.FormatMsg("Failed to receive request body"), bodyErr)
		errStr := bodyErr.Error()
		if code := strings.Split(errStr, "Code: "); len(code) > 0 {
			if "412" == code[1] {
				strList := make([]interface{}, 0)
				strList = append(strList, DelayOp)
				resultItemsChan <- ResultItems{Items: strList}
			} else {
				resultItemsChan <- NilResult()
			}
			return
		}
	}

	result := request.Parser(body)
	resultItemsChan <- result
	//fmt.Println("工作协程结束, 当前活跃线程数:\t", runtime.NumGoroutine())
}
