package model

import (
	"time"
)

type Shop struct {
	ShopId     int       `xorm:"not null pk autoincr INT(11)"`
	WxOpenId   string    `xorm:"unique VARCHAR(255)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"TIMESTAMP"`
	Phone      string    `xorm:"VARCHAR(11)"`
}
