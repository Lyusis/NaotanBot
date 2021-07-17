package server

import (
	"fmt"
	"net/http"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/service/cq/httpserver"
	"github.com/Lyusis/NaotanBot/service/cq/wsserver"
)

func HttpCQServer() {
	http.HandleFunc("/", httpserver.Handler)
	NewServer()
}

func WSCQServer() {
	http.HandleFunc("/api", wsserver.WSApiHandler)
	http.HandleFunc("/event", wsserver.WSEventHandler)
	NewServer()
}

func NewServer() {
	addr := fmt.Sprintf("%s:%d", conf.CQReceiver.IP, conf.CQReceiver.Port)
	go func(addr string) {
		serverErr := http.ListenAndServe(addr, nil)
		if serverErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("The listening server failed to start"), logger.FormatError(serverErr))
			return
		}
	}(addr)
	logger.Sugar.Info(logger.FormatMsg("Start listening server"), logger.FormatTitle("IP地址"), conf.CQSendDest.IP, logger.FormatTitle("端口"), conf.CQSendDest.Port)
}
