package model

import (
	"time"
)

type UserScoreLog struct {
	LogoId     int       `xorm:"not null pk autoincr INT(11)"`
	UserId     int       `xorm:"not null index INT(11)"`
	OrderId    int       `xorm:"not null index INT(11)"`
	Logo       string    `xorm:"VARCHAR(100)"`
	Score      int       `xorm:"not null default 0 INT(11)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
