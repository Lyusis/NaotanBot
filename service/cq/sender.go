package cq

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Lyusis/NaotanBot/conf"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/utils"
)

// WSSender websocket
type WSSender struct{}

func (sender *WSSender) SendGroupMessage(groupId int, message string) {
	msgBox := MsgBox{
		Action: SendGroup, Id: groupId, Message: message,
	}
	MsgBoxChan <- msgBox
}

func (sender *WSSender) SendPrivateMessage(userId int, message string) {
	msgBox := MsgBox{
		Action: SendPrivate, Id: userId, Message: message,
	}
	MsgBoxChan <- msgBox
}

// HttpSender http
type HttpSender struct{}

func (sender *HttpSender) SendGroupMessage(groupId int, message string) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	urlStr := fmt.Sprintf("http://%s:%d/send_group_msg?group_id=%d&message=%s", conf.CQSendDest.IP, conf.CQSendDest.Port, groupId, message)
	if !strings.EqualFold("", conf.Token) {
		urlStr += "&access_token=" + conf.Token
	}
	logger.Sugar.Info("发送Q群消息", logger.FormatTitle("URL"), urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatError(requestErr))
		return
	}
	utils.BasicReceiver(client.Do(request))
}

func (sender *HttpSender) SendPrivateMessage(userId int, message string) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	urlStr := fmt.Sprintf("http://%s:%d/send_private_msg?user_id=%s&message=%s", conf.CQSendDest.IP, conf.CQSendDest.Port, userId, message)
	if !strings.EqualFold("", conf.Token) {
		urlStr += "&access_token=" + conf.Token
	}
	logger.Sugar.Info("发送私聊消息", logger.FormatTitle("URL"), urlStr)
	request, requestErr := http.NewRequest("GET", urlStr, nil)
	if requestErr != nil {
		logger.Sugar.Warn("发送消息失败", logger.FormatError(requestErr))
		return
	}
	utils.BasicReceiver(client.Do(request))
}
