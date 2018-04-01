package model

import (
	"time"
)

type PhotoPrintQueue struct {
	PrintQueueId int       `xorm:"not null pk autoincr INT(11)"`
	UserId       int       `xorm:"not null index INT(11)"`
	ShopId       int       `xorm:"not null index INT(11)"`
	UserPhotoId  int       `xorm:"not null index INT(11)"`
	PrintStatus  int       `xorm:"default 0 TINYINT(4)"`
	CreateTime   time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime   time.Time `xorm:"TIMESTAMP"`
}
