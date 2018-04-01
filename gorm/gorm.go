// Package gorm provides ...
package gormpool

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db, err = gorm.Open("mysql", "root:lsf000000@/cxj?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
