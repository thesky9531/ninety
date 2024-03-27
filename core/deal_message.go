package core

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"ninety/global"
	"ninety/model"
)

// 处理 cancel 函数
func UserCancel(clientMes model.Message) {
	switch {
	case clientMes.MatchType == "mmf" || clientMes.MatchType == "fmm":
		if _, ok := global.DRooms[clientMes.RoomID]; ok {
			delete(global.DRooms, clientMes.RoomID)
		}
	case clientMes.MatchType == "mmm" || clientMes.MatchType == "fmf":
		if _, ok := global.SRooms[clientMes.RoomID]; ok {
			delete(global.SRooms, clientMes.RoomID)
		}

	}
}

// 处理 join 函数
func UserJoin(clientMes model.Message, conn *websocket.Conn) {
	fmt.Println("handleUserJoin...")
	//实例client
	client := model.Client{
		UserID:      clientMes.UserID,
		UserGender:  clientMes.UserGender,
		MatchGender: clientMes.MatchGender,
		Conn:        conn,
	}

	//判断匹配类型
	if client.UserGender == "male" && client.MatchGender == "female" {
		fmt.Println("male match female...")
		client.MatchType = "mmf"
	} else if client.UserGender == "male" && client.MatchGender == "male" {
		fmt.Println("male match male...")
		client.MatchType = "mmm"
	} else if client.UserGender == "female" && client.MatchGender == "male" {
		fmt.Println("female match male...")
		client.MatchType = "fmm"
	} else if client.UserGender == "female" && client.MatchGender == "female" {
		fmt.Println("female match female...")
		client.MatchType = "fmf"
	}

	//调用 ClientManage 函数，将client存入相应的map中
	ClientManage(&client)
	global.UserMap[conn] = client

	//根据matchtype确认用户所在房间人数，判断是否需要发送通知
	if client.MatchType == "mmf" || client.MatchType == "fmm" {
		if len(global.DRooms[client.RoomID]) == 1 {
			//返回创建房间成功信息
			fmt.Println("==1,返回创建房间成功信息")
			creatRoomSuccess := model.Message{
				MesType:   model.SIGNAL_TYPE_CREATROOM,
				MatchType: client.MatchType,
				RoomID:    client.RoomID,
			}
			mesData, _ := json.Marshal(creatRoomSuccess)
			client.Conn.WriteMessage(1, mesData)
			return
		} else if len(global.DRooms[client.RoomID]) == 2 {
			//通知对方
			fmt.Println("==2,通知对方")
			//websocket.TextMessage: 表示消息是文本类型。
			//websocket.BinaryMessage: 表示消息是二进制类型。
			//websocket.CloseMessage: 表示关闭连接的消息。
			//websocket.PingMessage: 表示Ping消息。
			//websocket.PongMessage: 表示Pong消息。
			toFirstUserMes := model.Message{
				MesType:      model.SIGNAL_TYPE_NEW_PEER,
				RemoteUserID: client.UserID,
				MatchType:    client.MatchType,
				RoomID:       client.RoomID,
			}
			mesData, _ := json.Marshal(toFirstUserMes)
			for i := 0; i < len(global.DRooms[client.RoomID]); i++ {
				if global.DRooms[client.RoomID][i].Conn == conn {
					continue
				}
				global.DRooms[client.RoomID][i].Conn.WriteMessage(1, mesData)
			}
			return
		}
	} else if client.MatchType == "mmm" || client.MatchType == "fmf" {
		if len(global.SRooms[client.RoomID]) == 1 {
			//返回创建房间成功信息
			fmt.Println("==1,返回创建房间成功信息")
			creatRoomSuccess := model.Message{
				MesType:   model.SIGNAL_TYPE_CREATROOM,
				MatchType: client.MatchType,
				RoomID:    client.RoomID,
			}
			mesData, _ := json.Marshal(creatRoomSuccess)
			client.Conn.WriteMessage(1, mesData)
			return
		} else if len(global.SRooms[client.RoomID]) == 2 {
			//通知对方
			fmt.Println("==2,通知对方")
			//websocket.TextMessage: 表示消息是文本类型。
			//websocket.BinaryMessage: 表示消息是二进制类型。
			//websocket.CloseMessage: 表示关闭连接的消息。
			//websocket.PingMessage: 表示Ping消息。
			//websocket.PongMessage: 表示Pong消息。
			toFirstUserMes := model.Message{
				MesType:      model.SIGNAL_TYPE_NEW_PEER,
				RemoteUserID: client.UserID,
				MatchType:    client.MatchType,
				RoomID:       client.RoomID,
			}
			mesData, _ := json.Marshal(toFirstUserMes)
			for i := 0; i < len(global.SRooms[client.RoomID]); i++ {
				if global.SRooms[client.RoomID][i].Conn == conn {
					continue
				}
				global.SRooms[client.RoomID][i].Conn.WriteMessage(1, mesData)
			}
			return
		}
	}
}

// 处理 offer 转发函数
func TransOffer(clientMes model.Message, conn *websocket.Conn) {
	fmt.Println("handleTransOffer...")
	//打包resp-join
	toSecUserMes := model.Message{
		MesType:      model.SIGNAL_TYPE_RESP_JOIN,
		RemoteUserID: clientMes.UserID,
		MatchType:    clientMes.MatchType,
		RoomID:       clientMes.RoomID,
		Data:         clientMes.Data,
	}
	mesData, _ := json.Marshal(toSecUserMes)

	//通过MatchType和RoomID，找到对端conn
	if clientMes.MatchType == "mmf" || clientMes.MatchType == "fmm" {
		for i := 0; i < len(global.DRooms[clientMes.RoomID]); i++ {
			if global.DRooms[clientMes.RoomID][i].Conn == conn {
				continue
			}
			global.DRooms[clientMes.RoomID][i].Conn.WriteMessage(1, mesData)
		}
		return
	} else if clientMes.MatchType == "mmm" || clientMes.MatchType == "fmf" {
		for i := 0; i < len(global.SRooms[clientMes.RoomID]); i++ {
			if global.SRooms[clientMes.RoomID][i].Conn == conn {
				continue
			}
			global.SRooms[clientMes.RoomID][i].Conn.WriteMessage(1, mesData)
		}
		return
	}
}

// 处理 answer 转发函数
func TransAnswer(clientMes model.Message, conn *websocket.Conn) {
	fmt.Println("handleTransAnswer...")
	//打包answer
	toFirstUserMes := model.Message{
		MesType: model.SIGNAL_TYPE_ANSWER,
		Data:    clientMes.Data,
	}
	mesData, _ := json.Marshal(toFirstUserMes)

	//通过MatchType和RoomID，找到对端conn
	if clientMes.MatchType == "mmf" || clientMes.MatchType == "fmm" {
		for i := 0; i < len(global.DRooms[clientMes.RoomID]); i++ {
			if global.DRooms[clientMes.RoomID][i].Conn == conn {
				continue
			}
			global.DRooms[clientMes.RoomID][i].Conn.WriteMessage(1, mesData)
		}
		return
	} else if clientMes.MatchType == "mmm" || clientMes.MatchType == "fmf" {
		for i := 0; i < len(global.SRooms[clientMes.RoomID]); i++ {
			if global.SRooms[clientMes.RoomID][i].Conn == conn {
				continue
			}
			global.SRooms[clientMes.RoomID][i].Conn.WriteMessage(1, mesData)
		}
		return
	}
}

// 处理 candidate 转发函数
func TransCandidate(clientMes model.Message, conn *websocket.Conn) {
	fmt.Println("handleTransCandidate...")
	//打包 candidate
	toAnotherUserMes := model.Message{
		MesType: model.SIGNAL_TYPE_CANDIDATE,
		Data:    clientMes.Data,
	}
	mesData, _ := json.Marshal(toAnotherUserMes)

	//通过MatchType和RoomID，找到对端conn
	if clientMes.MatchType == "mmf" || clientMes.MatchType == "fmm" {
		for i := 0; i < len(global.DRooms[clientMes.RoomID]); i++ {
			if global.DRooms[clientMes.RoomID][i].Conn == conn {
				continue
			}
			global.DRooms[clientMes.RoomID][i].Conn.WriteMessage(1, mesData)
		}
		return
	} else if clientMes.MatchType == "mmm" || clientMes.MatchType == "fmf" {
		for i := 0; i < len(global.SRooms[clientMes.RoomID]); i++ {
			if global.SRooms[clientMes.RoomID][i].Conn == conn {
				continue
			}
			global.SRooms[clientMes.RoomID][i].Conn.WriteMessage(1, mesData)
		}
		return
	}
}
