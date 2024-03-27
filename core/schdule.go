package core

import (
	"fmt"
	"ninety/global"
	"ninety/model"
)

// 调度用户函数，用于将用户放到目标房间
func ClientManage(client *model.Client) {
	//根据matchtype将用户放对应的map中的房间中
	if client.MatchType == "mmf" {
		//mmf进入 dRooms
		fmt.Println("to dRooms...")
		//调用 判断进入/创建房间函数
		EnterOrCreatRoom(global.DRooms, &global.DRoomNum, "fmm", client)
		return
	} else if client.MatchType == "fmm" {
		//fmm进入 dRooms
		fmt.Println("to dRooms...")
		//调用 判断进入/创建房间函数
		EnterOrCreatRoom(global.DRooms, &global.DRoomNum, "mmf", client)
		return
	} else if client.MatchType == "mmm" {
		//mmm 进入 sRooms
		fmt.Println("to sRooms...")
		//调用 判断进入/创建房间函数
		EnterOrCreatRoom(global.SRooms, &global.SRoomNum, "mmm", client)
		return
	} else if client.MatchType == "fmf" {
		//fmf 进入 sRooms
		fmt.Println("to sRooms...")
		//调用 判断进入/创建房间函数
		EnterOrCreatRoom(global.SRooms, &global.SRoomNum, "fmf", client)
		return
	}
}

// 判断进入/创建房间函数
func EnterOrCreatRoom(roomsTp map[int][]model.Client, roomNum *int, findMatchTp string, client *model.Client) {
	//如果房间数量等于0，则创建房间，并将用户放进房间
	if len(roomsTp) == 0 {
		*roomNum++
		client.RoomID = *roomNum
		roomsTp[*roomNum] = []model.Client{}
		roomsTp[*roomNum] = append(roomsTp[*roomNum], *client)
		fmt.Println("noroom end creatroom...")
		return
	}
	//如果房间数量不等于0，则寻找用户数量为1且房间等待用户为findMatchTp的房间，并将用户放进房间
	for key, value := range roomsTp {
		if len(value) == 1 && value[0].MatchType == findMatchTp {
			client.RoomID = key
			roomsTp[key] = append(roomsTp[key], *client)
			fmt.Println("for and enter...")
			return
		}
	}
	//如果没有，则创建房间，并将用户放进房间
	fmt.Println("for end creatroom...")
	*roomNum++
	client.RoomID = *roomNum
	roomsTp[*roomNum] = []model.Client{}
	roomsTp[*roomNum] = append(roomsTp[*roomNum], *client)
}
