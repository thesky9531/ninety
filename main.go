package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"ninety/controller"
)

func main() {
	// 解析命令行参数
	var port string
	flag.StringVar(&port, "port", "8888", "HTTP service port")
	flag.Parse()

	r := gin.Default()
	r.GET("/ws", controller.DealWss)
	r.Run(":" + port)
}
