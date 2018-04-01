// Package routes provides ...
package routes

import (
	//"cafe.lsfoo.com/controllers"
	"cafe.lsfoo.com/controllers/photo"
	"cafe.lsfoo.com/controllers/print"
	"cafe.lsfoo.com/controllers/shop"
	//"cafe.lsfoo.com/controllers/shopgoodscategory"
	"cafe.lsfoo.com/controllers/shopinfo"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func RouterStart() {

	r := mux.NewRouter()
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	//商店
	r.HandleFunc("/api/shop", shop.NewHandler).Methods("POST")
	r.HandleFunc("/api/shop", shop.FromWxCodeHandler).Methods("GET")

	r.HandleFunc("/api/ruler/shop", shop.FromRulerHandler).Methods("GET")

	//商店头像
	r.HandleFunc("/api/shop-avatar", shop.UpdateAvatarHandler).Methods("POST")

	r.HandleFunc("/api/shop-carousel/{id}", printq.FromCarouselHandler).Methods("GET")
	r.HandleFunc("/api/shop-position", shopinfo.UpdatePositionHandler).Methods("PUT")

	//商店详情
	r.HandleFunc("/api/shop-info", shopinfo.UpdateHandler).Methods("PUT")

	r.HandleFunc("/api/shop-info", shopinfo.NewHandler).Methods("POST")
	r.HandleFunc("/api/shop-info", shopinfo.FromGpsHandler).Methods("GET")
	r.HandleFunc("/api/shop-info/{id}", shopinfo.FromIdHandler).Methods("GET")
	r.HandleFunc("/api/shop-ruler-qr", shopinfo.FromRulerQrhandler).Methods("GET")
	r.HandleFunc("/api/shop-new-ruler", shopinfo.NewRulerHandler).Methods("POST")

	//图片
	r.HandleFunc("/api/user-photo", photo.NewHandler).Methods("POST")
	r.HandleFunc("/api/user-photo", photo.FromUserHandler).Methods("GET")
	r.HandleFunc("/api/user-photo/{id}", photo.DeleteHandler).Methods("DELETE")

	//打印列表
	r.HandleFunc("/api/print-queue", printq.FromUserHandler).Methods("GET")
	r.HandleFunc("/api/print-queue/{id}", printq.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/shop/printq", printq.ShopPrintQueueHandler).Methods("GET")
	r.HandleFunc("/shop/printq", printq.UpdateHandler).Methods("PUT")
	r.HandleFunc("/api/shop-printq", printq.FromShopHandler).Methods("GET")

	//店内商品分类

	//	r.HandleFunc("api/shop-goods-category", shopgoodscategory.NewHandler).Methods("POST")
	//	r.HandleFunc("api/shop-goods-category", shopgoodscategory.DeleteHandler).Methods("Delete")
	//	r.HandleFunc("api/shop-goods-category", shopgoodscategory.UpdateHandler).Methods("PUT")
	//	r.HandleFunc("api/shop-goods-category/{pid}", shopgoodscategory.FindHandler).Methods("GET")

	//天气
	//	r.HandleFunc("/api/weather", controllers.GetWeatherByIpHandler).Methods("GET")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err

		}
		// p will contain regular expression is compatible with regular expression in Perl, Python, and other languages.
		// for instance the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'
		p, err := route.GetPathRegexp()
		if err != nil {
			return err

		}
		m, err := route.GetMethods()
		if err != nil {
			return err

		}
		fmt.Println(strings.Join(m, ","), t, p)
		return nil

	})
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":6002", nil))

}
