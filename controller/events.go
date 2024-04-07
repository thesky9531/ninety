package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ninety/model"
	"ninety/service"
)

// CreateEvent 埋点请求
func CreateEvent(c *gin.Context) {
	var eventReq model.UserEvent
	if err := c.ShouldBind(&eventReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	if err := service.CreateEvent(&eventReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
