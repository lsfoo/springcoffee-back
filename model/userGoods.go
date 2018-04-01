package model

import (
	"time"
)

type UserGoods struct {
	UserGoodsId int       `xorm:"not null pk autoincr INT(11)"`
	UserId      int       `xorm:"not null index INT(11)"`
	IsSale      int       `xorm:"default 1 TINYINT(1)"`
	Name        string    `xorm:"not null VARCHAR(100)"`
	Description string    `xorm:"not null VARCHAR(500)"`
	Hot         int       `xorm:"TINYINT(4)"`
	Top         int       `xorm:"TINYINT(4)"`
	Rank        string    `xorm:"DECIMAL(10)"`
	Rec         int       `xorm:"TINYINT(4)"`
	Price       string    `xorm:"DECIMAL(10)"`
	IsUsed      string    `xorm:"VARCHAR(45)"`
	CreateTime  time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime  time.Time `xorm:"TIMESTAMP"`
}
