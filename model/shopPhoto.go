package model

type ShopPhoto struct {
	ShopId int    `xorm:"not null index INT(11)"`
	Rank   int    `xorm:"not null default 0 SMALLINT(6)"`
	Src    string `xorm:"not null VARCHAR(100)"`
}
