package printq

import (
	"cafe.lsfoo.com/auth"
	"cafe.lsfoo.com/controllers/user"
	"cafe.lsfoo.com/db"
	"cafe.lsfoo.com/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"html/template"
	"time"
)

var orm = db.Orm

type T struct {
	model.ShopInfo        `xorm:"extends"`
	model.UserPhoto       `xorm:"extends"`
	model.PhotoPrintQueue `xorm:"extends"`
}
type W struct {
	model.UserPhoto       `xorm:"extends"`
	model.PhotoPrintQueue `xorm:"extends"`
	model.UserInfo        `xorm:"extends"`
	CreateTime            time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

func FromUserHandler(w http.ResponseWriter, r *http.Request) {

	u, err := auth.FromUserAuth(r)
	if err != nil {
		return
	}
	results := make([]T, 0)
	//var photoPrintQueues []model.PhotoPrintQueue
	orm.Table("photo_print_queue").
		Desc("photo_print_queue.create_time").
		Join("LEFT", "shop_info", "photo_print_queue.shop_id = shop_info.shop_id").
		Join("LEFT", "user_photo", " photo_print_queue.user_photo_id = user_photo.user_photo_id ").
		Where("photo_print_queue.user_id = ?", u.UserId).Find(&results)
	jsons, _ := json.Marshal(results)
	fmt.Fprint(w, string(jsons))
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	_, err := user.FromAuthorization(r)
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var print model.PhotoPrintQueue
	print.PrintQueueId = id

	affected, _ := orm.Delete(print)
	fmt.Fprint(w, affected)
}
func ShopPrintQueueHandler(w http.ResponseWriter, r *http.Request) {
	shopId := r.FormValue("shop_id")

	results := make([]W, 0)
	err := orm.Table("photo_print_queue").Desc("photo_print_queue.create_time").
		Join("LEFT", "shop_info", "photo_print_queue.shop_id = shop_info.shop_id").
		Join("LEFT", "user_photo", " photo_print_queue.user_photo_id = user_photo.user_photo_id ").
		Join("LEFT", "user_info", " user_photo.user_id  = user_info.user_id").
		Where("photo_print_queue.shop_id = ?", shopId).Find(&results)
	if err != nil {
		panic(err)
	}
	//	jsons, _ := json.Marshal(results)
	//	fmt.Fprint(w, string(jsons))
	t, err := template.ParseFiles("./views/printq.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "printq", results)
}
func FromShopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Print(r)

	shopId, err := strconv.Atoi(r.FormValue("shop_id"))
	fmt.Print(shopId)
	if shopId == 0 {
		shop, err := auth.FromShopAuth(r)
		if err != nil {
			return
		}
		shopId = shop.ShopId

		fmt.Print("boss")

	} else {
		fmt.Print("ruler")
	}

	results := make([]W, 0)
	err = orm.Table("photo_print_queue").Desc("photo_print_queue.create_time").
		Join("LEFT", "shop_info", "photo_print_queue.shop_id = shop_info.shop_id").
		Join("LEFT", "user_photo", " photo_print_queue.user_photo_id = user_photo.user_photo_id ").
		Join("LEFT", "user_info", " user_photo.user_id  = user_info.user_id").
		Where("photo_print_queue.shop_id = ?", shopId).Find(&results)
	if err != nil {
		panic(err)
	}
	//	jsons, _ := json.Marshal(results)
	//	fmt.Fprint(w, string(jsons))
	jsons, _ := json.Marshal(results)
	fmt.Fprint(w, string(jsons))
}
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	var printq model.PhotoPrintQueue
	json.Unmarshal([]byte(body), &printq)
	orm.Update(printq)
	fmt.Fprint(w, printq)
}
func FromCarouselHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	results := make([]W, 0)
	err := orm.Table("photo_print_queue").Desc("photo_print_queue.create_time").
		Join("LEFT", "shop_info", "photo_print_queue.shop_id = shop_info.shop_id").
		Join("LEFT", "user_photo", " photo_print_queue.user_photo_id = user_photo.user_photo_id ").
		Join("LEFT", "user_info", " user_photo.user_id  = user_info.user_id").
		Where("photo_print_queue.shop_id = ?", id).Find(&results)
	if err != nil {
		panic(err)
	}
	jsons, _ := json.Marshal(results)
	fmt.Print(string(jsons))
	fmt.Fprint(w, string(jsons))
}
