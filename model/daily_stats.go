package model

type DailyStats struct {
	DateStr               string  `db:"date_str" json:"dateStr"`             // 统计日期
	WebsiteVisits         int     `db:"website_visits" json:"websiteVisits"` // 网站访问量
	MaleFemaleMatches     int     `db:"male_female_matches" json:"maleFemaleMatches"`
	MaleMaleMatches       int     `db:"male_male_matches" json:"maleMaleMatches"`
	FemaleMaleMatches     int     `db:"female_male_matches" json:"femaleMaleMatches"`
	FemaleFemaleMatches   int     `db:"female_female_matches" json:"femaleFemaleMatches"`
	TotalMatches          int     `db:"total_matches" json:"totalMatches,omitempty"`
	MaleFemaleSuccesses   int     `db:"male_female_successes" json:"maleFemaleSuccesses"`
	MaleMaleSuccesses     int     `db:"male_male_successes" json:"maleMaleSuccesses"`
	FemaleMaleSuccesses   int     `db:"female_male_successes" json:"femaleMaleSuccesses"`
	FemaleFemaleSuccesses int     `db:"female_female_successes" json:"femaleFemaleSuccesses"`
	TotalSuccesses        int     `db:"total_successes" json:"totalSuccesses,omitempty"`
	MaxWaitTime           int     `db:"max_wait_time" json:"maxWaitTime,omitempty"`
	MinMatchTime          int     `db:"min_match_time" json:"minMatchTime,omitempty"`
	AvgMatchTime          float64 `db:"avg_match_time" json:"avgMatchTime,omitempty"`
	MaxChatTime           int     `db:"max_chat_time" json:"maxChatTime,omitempty"`
	MinChatTime           int     `db:"min_chat_time" json:"minChatTime,omitempty"`
	AvgChatTime           float64 `db:"avg_chat_time" json:"avgChatTime,omitempty"`
	WaitTime5Seconds      int     `gorm:"column:wait_time_5_seconds" json:"wait_time_5_seconds,omitempty"`
	WaitTime10Seconds     int     `gorm:"column:wait_time_10_seconds" json:"wait_time_10_seconds,omitempty"`
	WaitTime20Seconds     int     `gorm:"column:wait_time_20_seconds" json:"wait_time_20_seconds,omitempty"`
	WaitTime40Seconds     int     `gorm:"column:wait_time_40_seconds" json:"wait_time_40_seconds,omitempty"`
	WaitTime60Seconds     int     `gorm:"column:wait_time_60_seconds" json:"wait_time_60_seconds,omitempty"`
	WaitTime90Seconds     int     `gorm:"column:wait_time_90_seconds" json:"wait_time_90_seconds,omitempty"`
	ChatTime5Seconds      int     `gorm:"column:chat_time_5_seconds" json:"chat_time_5_seconds,omitempty"`
	ChatTime10Seconds     int     `gorm:"column:chat_time_10_seconds" json:"chat_time_10_seconds,omitempty"`
	ChatTime20Seconds     int     `gorm:"column:chat_time_20_seconds" json:"chat_time_20_seconds,omitempty"`
	ChatTime40Seconds     int     `gorm:"column:chat_time_40_seconds" json:"chat_time_40_seconds,omitempty"`
	ChatTime60Seconds     int     `gorm:"column:chat_time_60_seconds" json:"chat_time_60_seconds,omitempty"`
	ChatTime90Seconds     int     `gorm:"column:chat_time_90_seconds" json:"chat_time_90_seconds,omitempty"`
}
