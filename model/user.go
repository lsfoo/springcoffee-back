package model

import (
	"time"
)

type User struct {
	UserId     int       `xorm:"not null pk autoincr INT(11)"`
	Phone      string    `xorm:"not null VARCHAR(11)"`
	WxOpenId   string    `xorm:"not null unique VARCHAR(255)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
