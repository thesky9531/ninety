package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"ninety/core"
	"ninety/global"
	"ninety/model"
)

func DealWss(conn *websocket.Conn) {
	// todo: 避免同时操作map，需要加锁
	global.UserMap[conn] = model.Client{}
	//客户端退出，删除连接
	defer func() {
		if _, ok := global.UserMap[conn]; ok {
			delete(global.UserMap, conn)
		}
	}()
	// 读取消息
	for {
		if conn == nil {
			fmt.Println("GetMes err:conn is nil")
			return
		}
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			closeErr, ok := err.(*websocket.CloseError)
			// 判断是否是客户端主动关闭
			if ok {
				if closeErr.Code == websocket.CloseGoingAway {
					//查询该链接的房间号，如果结果不为0，则代表已创建或进入房间
					if global.UserMap[conn].RoomID == 0 {
						return
					}
					//已创建或进入房间
					//根据matchtype确认用户所在房间人数，判断是否需要发送通知
					switch {
					case global.UserMap[conn].MatchType == "mmf" || global.UserMap[conn].MatchType == "fmm":
						if len(global.DRooms[global.UserMap[conn].RoomID]) == 2 {
							//通知对方
							fmt.Println("==2,通知对方")
							toFirstUserMes := model.Message{
								MesType: model.SIGNAL_TYPE_PEER_LEAVE,
							}
							mesData, _ := json.Marshal(toFirstUserMes)
							for i := 0; i < len(global.DRooms[global.UserMap[conn].RoomID]); i++ {
								if global.DRooms[global.UserMap[conn].RoomID][i].Conn == conn {
									continue
								}
								global.DRooms[global.UserMap[conn].RoomID][i].Conn.WriteMessage(1, mesData)
								if _, ok := global.UserMap[conn]; ok {
									delete(global.UserMap, conn)
								}
							}
						}
					case global.UserMap[conn].MatchType == "mmm" || global.UserMap[conn].MatchType == "fmf":
						if len(global.SRooms[global.UserMap[conn].RoomID]) == 2 {
							//通知对方
							fmt.Println("==2,通知对方")
							toFirstUserMes := model.Message{
								MesType: model.SIGNAL_TYPE_PEER_LEAVE,
							}
							mesData, _ := json.Marshal(toFirstUserMes)
							for i := 0; i < len(global.SRooms[global.UserMap[conn].RoomID]); i++ {
								if global.SRooms[global.UserMap[conn].RoomID][i].Conn == conn {
									continue
								}
								global.SRooms[global.UserMap[conn].RoomID][i].Conn.WriteMessage(1, mesData)
								if _, ok := global.UserMap[conn]; ok {
									delete(global.UserMap, conn)
								}
							}
						}
					}
					//判断结束，删除房间
					switch {
					case global.UserMap[conn].MatchType == "mmf" || global.UserMap[conn].MatchType == "fmm":
						if _, ok := global.DRooms[global.UserMap[conn].RoomID]; ok {
							delete(global.DRooms, global.UserMap[conn].RoomID)
						}
					case global.UserMap[conn].MatchType == "mmm" || global.UserMap[conn].MatchType == "fmf":
						if _, ok := global.SRooms[global.UserMap[conn].RoomID]; ok {
							delete(global.SRooms, global.UserMap[conn].RoomID)
						}
					}

				}
			}
		}

		switch messageType {
		case websocket.TextMessage:
			fmt.Println("handle websocket.TextMessage...")
			//反序列化信息内容
			var clientMes model.Message
			err = json.Unmarshal(p, &clientMes)
			if err != nil {
				fmt.Println("json.Unmarshal(p,&clientMes) err...")
			}
			//调用处理信息函数
			core.HandleMes(clientMes, conn)
		}
	}
}
