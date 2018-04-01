package model

import (
	"time"
)

type UserPhoto struct {
	UserPhotoId int       `xorm:"not null pk autoincr INT(11)"`
	UserId      int       `xorm:"not null index INT(11)"`
	Src         string    `xorm:"not null unique VARCHAR(100)"`
	Rank        int       `xorm:"default 0 INT(11)"`
	CreateTime  time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime  time.Time `xorm:"TIMESTAMP"`
	Words       string    `xorm:"VARCHAR(100)"`
}
