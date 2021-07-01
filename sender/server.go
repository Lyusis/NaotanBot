package sender

import (
	"github.com/Lyusis/NaotanMonitor/conf"
	"net/http"

	"github.com/Lyusis/NaotanMonitor/logger"
	"github.com/Lyusis/NaotanMonitor/sender/cq"
)

func CQServer() {
	addr := conf.CQReceiver.IP + ":" + conf.CQReceiver.Port
	go func() {
		NewHttpServer(addr)
	}()
}

func NewHttpServer(addr string) {
	http.HandleFunc("/", handlerFunc)

	logger.Sugar.Info(logger.FormatMsg("Start listening sender"), logger.FormatTitle("IP地址"), addr, logger.FormatTitle("端口"), 9001)

	addr += ":9001"

	serverErr := http.ListenAndServe(addr, nil)
	if serverErr != nil {
		logger.Sugar.Error(logger.FormatMsg("The listening sender failed to start"), logger.FormatError(serverErr))
	}
}

func handlerFunc(_ http.ResponseWriter, r *http.Request) {
	cq.Handler(r)
}
