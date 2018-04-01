package model

import (
	"time"
)

type ShopInfo struct {
	ShopId     int       `xorm:"not null index INT(11)"`
	ShopName   string    `xorm:"not null VARCHAR(100)"`
	Address    string    `xorm:"not null VARCHAR(100)"`
	CellPhone  string    `xorm:"not null VARCHAR(11)"`
	FixedPhone string    `xorm:"not null VARCHAR(20)"`
	Gps        string    `xorm:"VARCHAR(100)"`
	IsOpen     int       `xorm:"default 0 TINYINT(1)"`
	Top        int       `xorm:"TINYINT(4)"`
	Hot        int       `xorm:"TINYINT(4)"`
	Rank       string    `xorm:"DECIMAL(10)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"TIMESTAMP"`
	Avatar     string    `xorm:"VARCHAR(50)"`
	Lng        string    `xorm:"VARCHAR(50)"`
	Lat        string    `xorm:"VARCHAR(50)"`
	Id         int       `xorm:"not null pk autoincr INT(10)"`
}
