package global

import (
	"github.com/gorilla/websocket"
	"ninety/model"
)

// 设置用户map，可通过用户连接找到用户信息
var UserMap = make(map[*websocket.Conn]model.Client)

// 设置房间map
var DRooms = make(map[int][]model.Client)
var SRooms = make(map[int][]model.Client)

// 设置房间数量，即roomid
var DRoomNum = 10000
var SRoomNum = 10000
