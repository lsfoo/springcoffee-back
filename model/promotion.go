package model

type Promotion struct {
	PromotionId         int    `xorm:"not null pk autoincr INT(11)"`
	ShopId              int    `xorm:"not null index INT(11)"`
	PromotionCategoryId int    `xorm:"not null index INT(11)"`
	Tag                 string `xorm:"VARCHAR(255)"`
}
