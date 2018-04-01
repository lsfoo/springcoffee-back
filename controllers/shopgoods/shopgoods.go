package shopgoods

import (
	"cafe.lsfoo.com/auth"
	"cafe.lsfoo.com/db"
	"cafe.lsfoo.com/model"
	"encoding/json"
	//	"errors"
	"io/ioutil"
	"net/http"
	//	"strings"
	"fmt"
	"time"
)

var orm = db.Orm

func NewHandler(w http.ResponseWriter, r *http.Request) {
	shop, err := auth.FromShopAuth(r)
	if err != nil {
		painc(err)
		return
	}

	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	var shopGoods model.ShopGoods
	err := json.Unmarshal(body, &shopGoods)
	if err != nil {
		panic(err)
	}
	shopGoods.ShopId = shop.ShopId
	jsons := db.New(shopGoods)
	fmt.Fprint(w, jsons)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	shop, err := auth.FromShopAuth(r)
	if err != nil {
		painc(err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	shopGoods := &model.ShopGoods{ShopGoodsId: id}
	affected := db.Delete(shopGoods)
	fmt.Fprint(w, affected)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	shop, err := auth.FromShopAuth(r)
	if err != nil {
		painc(err)
		return
	}

	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	var shopGoods model.ShopGoods
	err := json.Unmarshal(body, &shopGoods)
	if err != nil {
		panic(err)
	}
	jsons := db.Update(shopGoods)
	fmt.Fprint(w, jsons)
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	shop, err := auth.FromShopAuth(r)
	if err != nil {
		painc(err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["pid"])

	var shopGoods []model.ShopGoods
	orm.Desc("create_time").Where("shop_id", shop.ShopId).Find(&shopGoods)
	if err != nil {
		panic(err)
	}
	jsons, _ := json.Marshal(shopGoods)
	fmt.Fprint(w, string(jsons))

}
