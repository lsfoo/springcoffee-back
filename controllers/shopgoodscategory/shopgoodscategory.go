package shopgoodscategory

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
	var shopGoodsCagegory model.ShopGoodsCagegory
	err := json.Unmarshal(body, &shopGoodsCagegory)
	if err != nil {
		panic(err)
	}
	shopGoodsCagegory.ShopId = shop.ShopId
	jsons := db.New(shopGoodsCagegory)
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
	shopGoodsCagegory := &model.ShopGoodsCagegory{ShopGoodsCategoryId: id}
	affected := db.Delete(shopGoodsCagegory)
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
	var shopGoodsCagegory model.ShopGoodsCagegory
	err := json.Unmarshal(body, &shopGoodsCagegory)
	if err != nil {
		panic(err)
	}
	shopGoodsCagegory.ShopId = shop.ShopId
	jsons := db.Update(shopGoodsCagegory)
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

	var shopGoodsCagegory []model.ShopGoodsCagegory
	orm.Desc("create_time").Where("pid = ?", id).And("shop_id", shop.ShopId).Find(&shopGoodsCagegory)
	if err != nil {
		panic(err)
	}
	jsons, _ := json.Marshal(shopGoodsCagegory)
	fmt.Fprint(w, string(jsons))

}
