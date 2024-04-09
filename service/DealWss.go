package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"ninety/core"
	"ninety/global"
	"ninety/model"
	"time"
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
						//if _, ok := global.UserMap[conn]; ok {
						//	delete(global.UserMap, conn)
						//}
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
						//if _, ok := global.UserMap[conn]; ok {
						//	delete(global.UserMap, conn)
						//}
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
			userEvent := &model.UserEvent{}
			userEvent.EventType = 4
			userEvent.UserId = global.UserMap[conn].UserID
			userEvent.EventTime.Time = time.Now()
			userEvent.DeviceId = "unknown"
			userEvent.EventProperties = "user leave"
			userEvent.Duration = time.Now().Sub(global.UserMap[conn].MatchDate).Milliseconds()
			userEvent.ActivityId = 1
			CreateEvent(userEvent)
			fmt.Println("conn.ReadMessage err:", err)
			return
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
			HandleMes(clientMes, conn)
		}
	}
}

// 处理信息函数
func HandleMes(clientMes model.Message, conn *websocket.Conn) {
	fmt.Println("HandleMes...")
	//根据mesType进入下一步流程
	switch clientMes.MesType {
	case model.SIGNAL_TYPE_JOIN:
		fmt.Println("MesType-join...")
		//调用userJoin函数
		core.UserJoin(clientMes, conn)
		// 开始匹配
		global.UserMapMutex.Lock()
		if client, exist := global.UserMap[conn]; exist {
			client.MatchDate = time.Now()
		}
		global.UserMapMutex.Unlock()
		userEvent := &model.UserEvent{}
		userEvent.EventType = 2
		userEvent.UserId = global.UserMap[conn].UserID
		userEvent.EventTime.Time = time.Now()
		userEvent.DeviceId = "unknown"
		userEvent.EventProperties = "user join"
		userEvent.Duration = 0
		userEvent.ActivityId = 1
		userEvent.UserGender = global.UserMap[conn].UserGender
		userEvent.MatchGender = global.UserMap[conn].MatchGender
		CreateEvent(userEvent)
	case model.SIGNAL_TYPE_CANCEL:
		fmt.Println("MesType-cancel...")
		//调用userCancel函数
		core.UserCancel(clientMes)
	case model.SIGNAL_TYPE_LEAVE:
		fmt.Println("MesType-leave...")
		//调用userLeave函数
		//userLeave(mes, conn)
	case model.SIGNAL_TYPE_OFFER:
		fmt.Println("MesType-offer...")
		//调用tansoffer函数
		core.TransOffer(clientMes, conn)
		// 匹配成功
		global.UserMapMutex.Lock()
		if client, exist := global.UserMap[conn]; exist {
			client.ChatDate = time.Now()
			client.MatchTime = client.ChatDate.Sub(client.MatchDate).Milliseconds()
		}
		global.UserMapMutex.Unlock()
		userEvent := &model.UserEvent{}
		userEvent.EventType = 3
		userEvent.UserId = global.UserMap[conn].UserID
		userEvent.EventTime.Time = time.Now()
		userEvent.DeviceId = "unknown"
		userEvent.EventProperties = "match success"
		userEvent.Duration = global.UserMap[conn].MatchTime
		userEvent.ActivityId = 1
		userEvent.UserGender = global.UserMap[conn].UserGender
		userEvent.MatchGender = global.UserMap[conn].MatchGender
		CreateEvent(userEvent)
	case model.SIGNAL_TYPE_ANSWER:
		fmt.Println("MesType-answer...")
		//调用tansanswer函数
		core.TransAnswer(clientMes, conn)
		// 匹配成功
		global.UserMapMutex.Lock()
		if client, exist := global.UserMap[conn]; exist {
			client.ChatDate = time.Now()
			client.MatchTime = client.ChatDate.Sub(client.MatchDate).Milliseconds()
		}
		global.UserMapMutex.Unlock()
		userEvent := &model.UserEvent{}
		userEvent.EventType = 3
		userEvent.UserId = global.UserMap[conn].UserID
		userEvent.EventTime.Time = time.Now()
		userEvent.DeviceId = "unknown"
		userEvent.EventProperties = "match success"
		userEvent.Duration = global.UserMap[conn].MatchTime
		userEvent.ActivityId = 1
		userEvent.UserGender = global.UserMap[conn].UserGender
		userEvent.MatchGender = global.UserMap[conn].MatchGender
		CreateEvent(userEvent)
	case model.SIGNAL_TYPE_CANDIDATE:
		fmt.Println("MesType-candidate...")
		//调用tanscandidate函数
		core.TransCandidate(clientMes, conn)
	default:
		fmt.Println("MesType-unknow...")
		return
	}
}
