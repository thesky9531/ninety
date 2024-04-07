package dao

import (
	"ninety/common"
	"ninety/model"
)

func CreateEvent(eventReq *model.UserEvent) error {

	return common.DB.Create(&eventReq).Error
}
