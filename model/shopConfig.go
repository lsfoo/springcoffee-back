package model

import (
	"time"
)

type ShopConfig struct {
	ShopConfigId      int       `xorm:"not null pk autoincr INT(11)"`
	ShopShopId        int       `xorm:"not null index INT(11)"`
	PrintPhotoMoney   string    `xorm:"DECIMAL(10)"`
	PrintPhotoScore   int       `xorm:"INT(11)"`
	CreateTime        time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime        time.Time `xorm:"TIMESTAMP"`
	ScoreRatio        string    `xorm:"default 0 DECIMAL(10)"`
	ScoreToMoneyRatio string    `xorm:"default 0 DECIMAL(10)"`
	MoneyToScoreRatio string    `xorm:"default 0 DECIMAL(10)"`
}
