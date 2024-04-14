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
		dailyStat.TotalMatches++
		switch eventReq.UserGender {
		case "female":
			switch eventReq.MatchGender {
			case "female":
				dailyStat.FemaleFemaleMatches++
			case "male":
				dailyStat.FemaleMaleMatches++
			}
		case "male":
			switch eventReq.MatchGender {
			case "female":
				dailyStat.MaleFemaleMatches++
			case "male":
				dailyStat.MaleMaleMatches++
			}
		}
	case 3:
		dailyStat.TotalSuccesses++
		switch eventReq.UserGender {
		case "female":
			switch eventReq.MatchGender {
			case "female":
				dailyStat.FemaleFemaleSuccesses++
			case "male":
				dailyStat.FemaleMaleSuccesses++
			}
		case "male":
			switch eventReq.MatchGender {
			case "female":
				dailyStat.MaleFemaleSuccesses++
			case "male":
				dailyStat.MaleMaleSuccesses++
			}
		}
		dailyStat.WaitTimeSum += eventReq.Duration
		if int(eventReq.Duration) > dailyStat.MaxWaitTime {
			dailyStat.MaxWaitTime = int(eventReq.Duration)
		}
		if int(eventReq.Duration) < dailyStat.MinMatchTime && int(eventReq.Duration) > 0 {
			dailyStat.MinMatchTime = int(eventReq.Duration)
		}
		dailyStat.AvgMatchTime = float64(dailyStat.WaitTimeSum) / float64(dailyStat.TotalSuccesses)
		if eventReq.Duration > 5 {
			dailyStat.WaitTime5Seconds++
		}
		if eventReq.Duration > 10 {
			dailyStat.WaitTime10Seconds++
		}
		if eventReq.Duration > 20 {
			dailyStat.WaitTime20Seconds++
		}
		if eventReq.Duration > 40 {
			dailyStat.WaitTime40Seconds++
		}
		if eventReq.Duration > 60 {
			dailyStat.WaitTime60Seconds++
		}
		if eventReq.Duration > 80 {
			dailyStat.WaitTime90Seconds++
		}

	case 4:
		if eventReq.IsChat {
			dailyStat.ChatTimeSum += eventReq.Duration
			if int(eventReq.Duration) > dailyStat.MaxChatTime {
				dailyStat.MaxChatTime = int(eventReq.Duration)
			}
			if int(eventReq.Duration) < dailyStat.MinChatTime {
				dailyStat.MinChatTime = int(eventReq.Duration)
			}
			dailyStat.AvgChatTime = float64(dailyStat.ChatTimeSum) / float64(dailyStat.TotalSuccesses)
			if eventReq.Duration > 5 {
				dailyStat.ChatTime5Seconds++
			}
			if eventReq.Duration > 10 {
				dailyStat.ChatTime10Seconds++
			}
			if eventReq.Duration > 20 {
				dailyStat.ChatTime20Seconds++
			}
			if eventReq.Duration > 40 {
				dailyStat.ChatTime40Seconds++
			}
			if eventReq.Duration > 60 {
				dailyStat.ChatTime60Seconds++
			}
			if eventReq.Duration > 80 {
				dailyStat.ChatTime90Seconds++
			}
		} else {
			dailyStat.WaitTimeSum += eventReq.Duration
			dailyStat.AvgMatchTime = float64(dailyStat.WaitTimeSum) / float64(dailyStat.TotalSuccesses)
			if int(eventReq.Duration) > dailyStat.MaxWaitTime {
				dailyStat.MaxWaitTime = int(eventReq.Duration)
			}
			if eventReq.Duration > 5 {
				dailyStat.WaitTime5Seconds++
			}
			if eventReq.Duration > 10 {
				dailyStat.WaitTime10Seconds++
			}
			if eventReq.Duration > 20 {
				dailyStat.WaitTime20Seconds++
			}
			if eventReq.Duration > 40 {
				dailyStat.WaitTime40Seconds++
			}
			if eventReq.Duration > 60 {
				dailyStat.WaitTime60Seconds++
			}
			if eventReq.Duration > 80 {
				dailyStat.WaitTime90Seconds++
			}
		}

	}

	if err := dao.UpdateDailyStat(dailyStat); err != nil {
		log.Println("failed to update daily stats: %w", err)
		err = errors.New("failed to update daily stats")
		return err
	}
	// 如果执行成功，返回 nil 表示没有错误
	return nil
}
