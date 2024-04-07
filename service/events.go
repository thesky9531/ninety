package service

import (
	"errors"
	"log"
	"ninety/dao"
	"ninety/model"
	"time"
)

func CreateEvent(eventReq *model.UserEvent) error {
	eventReq.EventTime.Time = time.Now()

	if err := dao.CreateEvent(eventReq); err != nil {
		log.Println("failed to execute SQL statement: %w", err)
		err = errors.New("failed to execute SQL statement")
		return err
	}
	// 获取当天的daily_stats记录,如果没有则创建
	dailyStat, err := dao.GetOrCreateDailyStat()
	if err != nil {
		log.Println("failed to get or create daily stats: %w", err)
		err = errors.New("failed to get or create daily stats")
		return err
	}
	switch eventReq.EventType {
	case 1:
		dailyStat.WebsiteVisits++
	case 2:

	}

	if err := dao.UpdateDailyStat(dailyStat); err != nil {
		log.Println("failed to update daily stats: %w", err)
		err = errors.New("failed to update daily stats")
		return err
	}
	// 如果执行成功，返回 nil 表示没有错误
	return nil
}
