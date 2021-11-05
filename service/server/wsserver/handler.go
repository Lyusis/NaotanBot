package wsserver

import (
	"encoding/json"
	"github.com/Lyusis/NaotanBot/logger"
	"github.com/Lyusis/NaotanBot/service/cq"
	"github.com/Lyusis/NaotanBot/service/server/register"
	"github.com/gorilla/websocket"
	"net/http"
)

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
				register.Register(message)
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
			case msgBox := <-cq.MsgBoxChan:
				var (
					marshaledMsg []byte
					err          error
				)
				switch msgBox.Action {
				case cq.SendGroup:
					marshaledMsg, err = groupMessage(msgBox.Message, msgBox.Id)
				case cq.SendPrivate:
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

func groupMessage(message string, groupId int) ([]byte, error) {
	sendMessage := cq.GroupMessage{
		Action: cq.SendGroup,
	}
	sendMessage.Params.Message = message
	sendMessage.Params.GroupId = groupId
	marshal, err := json.Marshal(sendMessage)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func privateMessage(message string, privateId int) ([]byte, error) {
	sendMessage := cq.IndividualMessage{
		Action: cq.SendPrivate,
	}
	sendMessage.Params.Message = message
	sendMessage.Params.UserId = privateId
	marshal, err := json.Marshal(sendMessage)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
