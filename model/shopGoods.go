package model

import (
	"time"
)

type ShopGoods struct {
	ShopGoodsId     int       `xorm:"not null pk autoincr INT(11)"`
	ShopId          int       `xorm:"not null index INT(11)"`
	GoodsCategoryId int       `xorm:"not null index INT(11)"`
	Price           string    `xorm:"DECIMAL(10)"`
	IsSale          int       `xorm:"default 1 TINYINT(1)"`
	Name            string    `xorm:"not null VARCHAR(100)"`
	Description     string    `xorm:"not null VARCHAR(500)"`
	Hot             int       `xorm:"TINYINT(4)"`
	Top             int       `xorm:"TINYINT(4)"`
	Rank            string    `xorm:"DECIMAL(10)"`
	Rec             int       `xorm:"TINYINT(4)"`
	CreateTime      time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime      time.Time `xorm:"TIMESTAMP"`
	Ratio           string    `xorm:"DECIMAL(10)"`
}
