package model

import (
	"time"
)

type Order struct {
	OrderId     int       `xorm:"not null pk autoincr INT(11)"`
	ShopId      int       `xorm:"not null index INT(11)"`
	UserId      int       `xorm:"not null index INT(11)"`
	OrderStatus int       `xorm:"not null default 0 SMALLINT(6)"`
	TotalFee    string    `xorm:"default 0 DECIMAL(10)"`
	CreateTime  time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime  time.Time `xorm:"TIMESTAMP"`
	Ordercol    string    `xorm:"VARCHAR(45)"`
}
