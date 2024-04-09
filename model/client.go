package model

import (
	"github.com/gorilla/websocket"
	"time"
)

// 前后端通讯消息
type Message struct {
	MesType      string      `json:"mesType"`
	UserID       string      `json:"userId"`
	RemoteUserID string      `json:"remoteUserID"`
	UserGender   string      `json:"userGender"`
	MatchGender  string      `json:"matchGender"`
	Data         interface{} `json:"data"`
	MatchType    string      `json:"matchType"`
	RoomID       int         `json:"roomID"`
}

// 设置用户信息结构体
type Client struct {
	UserID      string
	UserGender  string
	MatchGender string
	MatchType   string
	RoomID      int
	Conn        *websocket.Conn
	MatchDate   time.Time
	MatchTime   int64
	ChatDate    time.Time
	ChatTime    int64
}
