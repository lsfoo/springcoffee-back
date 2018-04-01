package model

type PromotionCategory struct {
	PromotionCategoryId int    `xorm:"not null pk INT(11)"`
	Name                string `xorm:"not null VARCHAR(255)"`
	IconSrc             string `xorm:"VARCHAR(45)"`
}
