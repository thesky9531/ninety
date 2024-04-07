package dao

import (
	"ninety/common"
	"ninety/model"
	"time"
)

func GetOrCreateDailyStat() (dailyStat *model.DailyStats, err error) {
	// 查询当天的daily_stats记录
	dailyStat = &model.DailyStats{}
	err = common.DB.Where("date_str = ?", time.Now().Format("2006-01-02")).First(dailyStat).Error
	if err != nil {
		if err.Error() == "record not found" {
			// 如果没有记录，则创建
			dailyStat = &model.DailyStats{
				DateStr: time.Now().Format("2006-01-02"),
			}
			err = common.DB.Create(dailyStat).Error
		} else {
			return
		}
	}
	return dailyStat, err
}

func UpdateDailyStat(dailyStat *model.DailyStats) error {
	return common.DB.Where("date_str = ?", time.Now().Format("2006-01-02")).Save(dailyStat).Error
}
