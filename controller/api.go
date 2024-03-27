package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
	"ninety/service"
)

func DealWss(c *gin.Context) {
	upGrade := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Error("deal wss error:", err)
	}
	defer conn.Close()
	service.DealWss(conn)
}
