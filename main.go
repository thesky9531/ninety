package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"ninety/common"
	"ninety/controller"
)

func main() {
	// 解析命令行参数
	var port string
	var dbDrive string
	flag.StringVar(&port, "port", "8888", "HTTP service port")
	flag.StringVar(&dbDrive, "db", "mysql", "database drive")
	flag.Parse()

	common.Init(dbDrive)
	r := gin.Default()
	r.GET("/ws", controller.DealWss)

	// 埋点接口
	r.POST("/api/v1/events", controller.CreateEvent)
	r.Run(":" + port)
}
