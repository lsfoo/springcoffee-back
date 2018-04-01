package model

type ShopGoodsPhoto struct {
	PhotoId int    `xorm:"not null pk autoincr INT(11)"`
	GoodsId int    `xorm:"not null index INT(11)"`
	Src     string `xorm:"not null VARCHAR(100)"`
	Rank    int    `xorm:"default 0 TINYINT(4)"`
}
