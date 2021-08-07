package wsserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/service/common"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/gorilla/websocket"
)

var SendTool = cq.Sender{
	SendMessage: &WSSender{},
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSEventHandler(w http.ResponseWriter, r *http.Request) {
	conn, connErr := upgrader.Upgrade(w, r, nil)
	if connErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("WebSocket Upgrade ERROR(Event)"), logger.FormatError(connErr))
		return
	}

	go func(conn *websocket.Conn) {
		eventMessage := cq.MetaEventMessage{}
		message := cq.MessageMessage{}
		postType := cq.PostType{}
		for {
			// 读取客户端的消息
			_, msg, readMessageErr := conn.ReadMessage()
			if readMessageErr != nil {
				logger.Sugar.Warn(logger.FormatMsg("WebSocket cannot read message"), logger.FormatError(readMessageErr))
				return
			}
			checkErr := json.Unmarshal(msg, &postType)
			if checkErr != nil {
				logger.Sugar.Warn("Server failed to parse JSON message(TYPE)", logger.FormatError(checkErr))
			}

			switch postType.PostType {
			case "meta_event":
				jsonErr := json.Unmarshal(msg, &eventMessage)
				if jsonErr != nil {
					logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(META_EVENT)"), logger.FormatError(jsonErr))
				}
			case "message":
				jsonErr := json.Unmarshal(msg, &message)
				if jsonErr != nil {
					logger.Sugar.Warn(logger.FormatMsg("Server failed to parse JSON message(MESSAGE)"), logger.FormatError(jsonErr))
				}
				SendTool.AJun(message)
				SendTool.InsertVup(message)
			}
		}
	}(conn)
}

func WSApiHandler(w http.ResponseWriter, r *http.Request) {
	conn, connErr := upgrader.Upgrade(w, r, nil)
	if connErr != nil {
		logger.Sugar.Warn(logger.FormatMsg("WebSocket Upgrade ERROR(Api)"), logger.FormatError(connErr))
		return
	}
	go func(conn *websocket.Conn) {
		for {
			select {
			case msgBox := <-common.MsgBoxChan:
				var (
					marshaledMsg []byte
					err          error
				)
				switch msgBox.Action {
				case common.SendGroup:
					marshaledMsg, err = groupMessage(msgBox.Message, msgBox.Id)
				case common.SendPrivate:
					marshaledMsg, err = privateMessage(msgBox.Message, msgBox.Id)
				}
				if err != nil {
					logger.Sugar.Warn(logger.FormatMsg("Format Message ERROR"), logger.FormatError(err))
					continue
				}
				writeMessageErr := conn.WriteMessage(websocket.TextMessage, marshaledMsg)
				if writeMessageErr != nil {
					logger.Sugar.Warn(logger.FormatMsg("WebSocket cannot read message"), logger.FormatError(writeMessageErr))
					return
				}
			}
		}
	}(conn)
}

func groupMessage(message string, groupId string) ([]byte, error) {
	sendMessage := cq.GroupMessage{
		Action: common.SendGroup,
	}
	sendMessage.Params.Message = message
	sendMessage.Params.GroupId, _ = strconv.Atoi(groupId)
	marshal, err := json.Marshal(sendMessage)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func privateMessage(message string, privateId string) ([]byte, error) {
	sendMessage := cq.IndividualMessage{
		Action: common.SendPrivate,
	}
	sendMessage.Params.Message = message
	sendMessage.Params.UserId, _ = strconv.Atoi(privateId)
	marshal, err := json.Marshal(sendMessage)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
