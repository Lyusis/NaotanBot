package api

import (
	"monitor/config"
	"net/http"
)

type saberserverResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendBarkMessage(title, desp string) {
	url := "https://sctapi.ftqq.com/" + config.SaberchanCode +".send?title=" + title + "&desp=" + desp
	header := "application/json;charset=UTF-8"
	BasicReceiver(http.Post(url, header, nil))
}