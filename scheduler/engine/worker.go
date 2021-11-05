package engine

import (
	"net/http"
	"strings"

	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/scheduler/fetcher"
)

func Worker(request Request, resultItemsChan chan ResultItems) {

	if request.Name != "" {
		logger.Sugar.Info(logger.FormatMsg("Fetching"), logger.FormatTitle("URL"), request.Url, logger.FormatTitle("MissionTitle"), request.Name)
	} else {
		logger.Sugar.Info(logger.FormatMsg("Fetching"), logger.FormatTitle("URL"), request.Url)
	}
	if request.Method == "" {
		request.Method = http.MethodGet
	}
	body, bodyErr := fetcher.Fetcher(request.Url, request.Method, request.Body)
	if bodyErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("Failed to receive request body"), bodyErr)
		errStr := bodyErr.Error()
		// TODO 重新设计处理方案, 封装给server
		if code := strings.Split(errStr, "Code: "); len(code) > 1 {
			if string(rune(http.StatusOK)) != code[1] {
				strList := make([]interface{}, 0)
				strList = append(strList, DelayOp)
				resultItemsChan <- ResultItems{Items: strList}
				//basic.SendDelayMsg()
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
