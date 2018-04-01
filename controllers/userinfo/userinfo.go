package userinfo

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
	user, err := auth.FromUserAuth(r)
	if err != nil {
		painc(err)
		return
	}

	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	var userInfo model.UserInfo
	err := json.Unmarshal(body, &userInfo)
	if err != nil {
		panic(err)
	}
	userInfo.ShopId = user.ShopId
	jsons := db.New(userInfo)
	fmt.Fprint(w, jsons)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	user, err := auth.FromUserAuth(r)
	if err != nil {
		painc(err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	userInfo := &model.UserInfo{UserInfoId: id}
	affected := db.Delete(userInfo)
	fmt.Fprint(w, affected)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.FromUserAuth(r)
	if err != nil {
		painc(err)
		return
	}

	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	var userInfo model.UserInfo
	err := json.Unmarshal(body, &userInfo)
	if err != nil {
		panic(err)
	}
	jsons := db.Update(userInfo)
	fmt.Fprint(w, jsons)
}

func GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.FromUserAuth(r)
	if err != nil {
		painc(err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["pid"])

	var userInfo []model.UserInfo
	orm.Desc("create_time").Where("shop_id", user.ShopId).Find(&userInfo)
	if err != nil {
		panic(err)
	}
	jsons, _ := json.Marshal(userInfo)
	fmt.Fprint(w, string(jsons))
}
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.FromUserAuth(r)
	if err != nil {
		painc(err)
		return
	}
	userInfo := &model.UserInfo{UserId: user.UserId}
	orm.Find(&userInfo)

	jsons, _ := json.Marshal(userInfo)
	fmt.Fprint(w, string(jsons))
}
