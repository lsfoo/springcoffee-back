package model

type ShopCategory struct {
	CategoryId int    `xorm:"not null pk INT(11)"`
	Pid        int    `xorm:"default 0 INT(11)"`
	Name       string `xorm:"not null VARCHAR(255)"`
	IsEnd      int    `xorm:"default 1 TINYINT(4)"`
	Rank       int    `xorm:"TINYINT(4)"`
	ShopId     int    `xorm:"not null index INT(11)"`
}
