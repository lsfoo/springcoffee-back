package model

type ShopGoodsCategory struct {
	ShopGoodsCategoryId int    `xorm:"not null pk autoincr INT(11)"`
	ShopId              int    `xorm:"not null index INT(11)"`
	Name                string `xorm:"not null VARCHAR(100)"`
	Pid                 int    `xorm:"default 0 INT(11)"`
	Rank                int    `xorm:"TINYINT(4)"`
	Description         string `xorm:"VARCHAR(45)"`
}
