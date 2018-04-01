package model

import (
	"time"
)

type Rqmt struct {
	RqmtId       int       `xorm:"not null pk autoincr INT(11)"`
	ShopGoodsId  int       `xorm:"not null index INT(11)"`
	Requirements string    `xorm:"not null VARCHAR(45)"`
	CreateTime   time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime   time.Time `xorm:"TIMESTAMP"`
	Rank         int       `xorm:"default 0 INT(11)"`
}
