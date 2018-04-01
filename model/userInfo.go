package model

type UserInfo struct {
	UserId   int    `xorm:"not null index INT(11)"`
	Phone    string `xorm:"VARCHAR(11)"`
	Sex      int    `xorm:"SMALLINT(6)"`
	Address  string `xorm:"VARCHAR(100)"`
	Score    int    `xorm:"default 0 INT(11)"`
	Name     string `xorm:"VARCHAR(45)"`
	NickName string `xorm:"VARCHAR(45)"`
}
