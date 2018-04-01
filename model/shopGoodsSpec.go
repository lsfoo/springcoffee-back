package model

type ShopGoodsSpec struct {
	ShopGoodsSpecId int    `xorm:"not null pk autoincr INT(11)"`
	ShopGoodsId     int    `xorm:"not null index INT(11)"`
	Price           string `xorm:"DECIMAL(10)"`
	PhotoSrc        string `xorm:"default '0' VARCHAR(100)"`
}
