package cq

import (
	"encoding/json"
	"fmt"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpHandler(_ http.ResponseWriter, r *http.Request) {
	eventMessage := MetaEventMessage{}
	message := MessageMessage{}
	postType := PostType{}
	readAll, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Sugar.Warn("Server failed to read JSON message", logger.FormatError(err))
	}
	checkErr := json.Unmarshal(readAll, &postType)
	if checkErr != nil {
		logger.Sugar.Warn("Server failed to parse JSON message(TYPE)", logger.FormatError(checkErr))
	}

	switch postType.PostType {
	case "meta_event":
		jsonErr := json.Unmarshal(readAll, &eventMessage)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(META_EVENT)"), logger.FormatError(jsonErr))
		}
	case "message":
		jsonErr := json.Unmarshal(readAll, &message)
		if jsonErr != nil {
			logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(MESSAGE)"), logger.FormatError(jsonErr))
		}
		AJun(message)
		At(message)
		AutoReturn(message)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(w http.ResponseWriter, r *http.Request) {

	conn, connErr := upgrader.Upgrade(w, r, nil)
	if connErr != nil {
		log.Println("Upgrade:", connErr)
		return
	}

	go func() {
		for {
			// 读取客户端的消息
			_, msg, readMessageErr := conn.ReadMessage()
			if readMessageErr != nil {
				return
			}

			fmt.Printf("Event: %+v\n", r.RemoteAddr)
			// 把消息打印到标准输出
			fmt.Printf("%s sent event: %s\n", conn.RemoteAddr(), string(msg))
		}
	}()
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	conn, connErr := upgrader.Upgrade(w, r, nil)
	if connErr != nil {
		log.Println("Upgrade:", connErr)
		return
	}

	SendTest(conn)

	go func() {
		for {
			// 读取客户端的消息
			_, msg, readMessageErr := conn.ReadMessage()
			if readMessageErr != nil {
				fmt.Printf("%s sent api: %s\n", conn.RemoteAddr(), string(msg))
				return
			}

			// 把消息打印到标准输出

			fmt.Printf("API: %+v\n", r.RemoteAddr)
		}
	}()

}

func SendTest(conn *websocket.Conn) {
	test1 := IndividualMessage{
		Action: "send_private_msg",
	}
	test1.Params.Message = "muamua"
	test1.Params.UserId = 506642268
	marshal, err := json.Marshal(test1)
	if err != nil {
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
