package main

import (
	"github.com/gin-gonic/gin"
	"ninety/controller"
)

func main() {
	r := gin.Default()
	r.GET("/ws", controller.DealWss)
	r.Run(":8888")
}
