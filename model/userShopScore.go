package model

type UserShopScore struct {
	UserId int `xorm:"not null index INT(11)"`
	ShopId int `xorm:"not null index INT(11)"`
	Score  int `xorm:"not null default 0 INT(11)"`
}
