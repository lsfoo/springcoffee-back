package db

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var Orm *xorm.Engine

func init() {
	var err error
	Orm, err = xorm.NewEngine("mysql", "root:lsf000000@/cxj?charset=utf8")
	Orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

}
func New(model interface{}) string {
	Orm.Insert(model)
	j, err := json.Marshal(&model)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func Update(model interface{}) string {
	Orm.Update(model)
	j, _ := json.Marshal(&model)
	return string(j)
}
func Delete(model interface{}) int64 {
	affected, _ := Orm.Delete(model)
	return affected
}
