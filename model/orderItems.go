package model

type OrderItems struct {
	OrderId     int    `xorm:"not null index INT(11)"`
	GoodsName   string `xorm:"VARCHAR(45)"`
	GoodsSpec   string `xorm:"VARCHAR(45)"`
	GoodsPrice  string `xorm:"DECIMAL(10)"`
	Quantity    int    `xorm:"INT(11)"`
	Price       string `xorm:"DECIMAL(10)"`
	GoodsSpecId int    `xorm:"not null index INT(11)"`
}
