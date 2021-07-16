package sender

import (
	"fmt"
	"github.com/Lyusis/NaotanMonitor/conf"
	cq2 "github.com/Lyusis/NaotanMonitor/service/cq"
	"net/http"

	"github.com/Lyusis/NaotanMonitor/logger"
)

func HttpCQServer() {
	NewServer(func(w http.ResponseWriter, r *http.Request) {
		cq2.HttpHandler(w, r)
	})
}

func WSCQServer() {
	NewServer(func(w http.ResponseWriter, r *http.Request) {
		cq2.WSHandler(w, r)
	})
}

func NewServer(_ func(http.ResponseWriter, *http.Request)) {

	addr := fmt.Sprintf("%s:%d", conf.CQReceiver.IP, conf.CQReceiver.Port)

	//http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/api", cq2.ApiHandler)
	http.HandleFunc("/event", cq2.WSHandler)
	go func(addr string) {
		serverErr := http.ListenAndServe(addr, nil)
		if serverErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("The listening sender failed to start"), logger.FormatError(serverErr))
			return
		}
	}(addr)
	logger.Sugar.Info(logger.FormatMsg("Start listening sender"), logger.FormatTitle("IP地址"), conf.CQSendDest.IP, logger.FormatTitle("端口"), conf.CQSendDest.Port)
}
