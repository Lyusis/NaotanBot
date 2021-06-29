package server

import (
	"net/http"

	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/Lyusis/NaotanMonitor/server/cq"
)

func NewHttpServer(addr string) {
	http.HandleFunc("/", handlerFunc)

	logger.Sugar.Info(logger.FormatMsg("Start listening server"), logger.FormatTitle("IP地址"), addr, logger.FormatTitle("端口"), 9001)

	addr += ":9001"

	serverErr := http.ListenAndServe(addr, nil)
	if serverErr != nil {
		logger.Sugar.Error(logger.FormatMsg("The listening server failed to start"), logger.FormatTitle("WRONG"), serverErr)
	}
}

func handlerFunc(_ http.ResponseWriter, r *http.Request) {
	cq.Handler(r)
}
