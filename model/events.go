package model

import (
	"database/sql"
)

type UserEvent struct {
	EventType       int64        `json:"event_type" db:"event_type"`
	UserId          string       `json:"user_id" db:"user_id"`
	DeviceId        string       `json:"device_id" db:"device_id"`
	EventProperties string       `json:"event_properties" db:"event_properties"`
	IpAddress       string       `json:"ip_address" db:"ip_address"`
	UserAgent       string       `json:"user_agent" db:"user_agent"`
	Country         string       `json:"country" db:"country"`
	City            string       `json:"city" db:"city"`
	ActivityId      int64        `json:"activity_id" db:"activity_id"`
	EventTime       sql.NullTime `json:"event_time" db:"event_time"`
	Duration        int64        `json:"duration" db:"duration"`
	UserGender      string       `json:"userGender" gorm:"-"`
	MatchGender     string       `json:"matchGender" gorm:"-"`
}
