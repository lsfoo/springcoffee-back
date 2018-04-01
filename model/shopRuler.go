package model

import (
	"time"
)

type ShopRuler struct {
	ShopRulerId int       `xorm:"not null pk autoincr INT(11)"`
	ShopId      int       `xorm:"INT(11)"`
	UserId      int       `xorm:"INT(11)"`
	CreateTime  time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime  time.Time `xorm:"TIMESTAMP"`
}
