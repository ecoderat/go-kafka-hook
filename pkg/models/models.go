package models

import "time"

type Logs struct {
	Host        string    `gorm:"not null"`
	Level       string    `gorm:"not null"`
	Msg         string    `gorm:"not null"`
	RequestType string    `gorm:"not null"`
	RespondTime int       `gorm:"not null" json:",string"`
	Time        time.Time `gorm:"not null"`
}
