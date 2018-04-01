package shop

import (
	"bufio"
	"cafe.lsfoo.com/auth"
	"cafe.lsfoo.com/db"
	"cafe.lsfoo.com/model"
	"encoding/json"
	//	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"

	"github.com/nfnt/resize"
	"image"
	"os"
	"strconv"
	"time"
	//	"strings"
	"fmt"

	"image/jpeg"
)

var orm = db.Orm

func FromWxCodeHandler(w http.ResponseWriter, r *http.Request) {
	shop, _ := auth.FromShopAuth(r)
	fmt.Print(shop.ShopId)

	shopInfo := model.ShopInfo{ShopId: shop.ShopId}
	if shop.ShopId > 0 {
		_, err := orm.Get(&shopInfo)
		if err != nil {
			panic(err)
		}
		jsons, _ := json.Marshal(shopInfo)
		fmt.Fprint(w, string(jsons))
		return
	}

	jsons, _ := json.Marshal(shopInfo)
	fmt.Fprint(w, string(jsons))

}
func NewHandler(w http.ResponseWriter, r *http.Request) {
	code, _ := auth.FromAuthHeader(r)
	wx, _ := auth.FromWXCode(code)
	shop := &model.Shop{WxOpenId: wx.OpenId, CreateTime: time.Now()}
	jsons := db.New(shop)
	if shop.ShopId > 0 {
		avatar, _, err := r.FormFile("file")
		if err != nil {
			panic(err)
			return
		}
		defer avatar.Close()

		uf, _, err := image.Decode(avatar)
		if err != nil {
			panic(err)
			return
		}

		canvasWidth := 800
		resizeImage := resize.Resize(uint(canvasWidth), 0, uf, resize.Lanczos3)

		fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".jpeg"
		filePath := "./public/upload/" + fileName

		outFile, err := os.Create(filePath)
		if err != nil {
			//log.Println(err)
			os.Exit(1)
		}

		defer outFile.Close()

		b := bufio.NewWriter(outFile)
		err = jpeg.Encode(b, resizeImage, nil)
		if err != nil {
			//log.Println(err)
			os.Exit(1)
		}
		err = b.Flush()
		if err != nil {
			//log.Println(err)
			os.Exit(1)
		}
		//fmt.Print(http.FileServer(http.Dir("./public")))
		client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAICmLY9c6bDHAX", "qUbMQOIx0qRRV11zSRHKR7E2rjuIUP")
		if err != nil {
			panic(err)
			return

		}

		bucket, err := client.Bucket("lfo")
		if err != nil {
			panic(err)
			return

		}

		err = bucket.PutObjectFromFile(fileName, filePath)
		if err != nil {
			panic(err)
			return
		}
		var shopInfo model.ShopInfo
		shopInfo.ShopId = shop.ShopId
		shopInfo.ShopName = r.FormValue("ShopName")
		shopInfo.Lng = r.FormValue("Lng")
		shopInfo.Lat = r.FormValue("Lat")
		shopInfo.Address = r.FormValue("Address")
		shopInfo.Avatar = fileName
		shopInfo.CreateTime = time.Now()
		orm.Insert(&shopInfo)
		j, _ := json.Marshal(shopInfo)

		fmt.Fprint(w, string(j))
		return
	}
	fmt.Fprint(w, jsons)
}
func UpdateAvatarHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	_, err = auth.FromShopAuth(r)
	if err != nil {
		return
	}
	shopId, _ := strconv.Atoi(r.FormValue("shopId"))
	avatar, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
		return
	}
	defer avatar.Close()

	uf, _, err := image.Decode(avatar)
	if err != nil {
		panic(err)
		return
	}

	canvasWidth := 400
	resizeImage := resize.Resize(uint(canvasWidth), 0, uf, resize.Lanczos3)

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".jpeg"
	filePath := "./public/upload/" + fileName

	outFile, err := os.Create(filePath)
	if err != nil {
		//log.Println(err)
		os.Exit(1)
	}

	defer outFile.Close()

	b := bufio.NewWriter(outFile)
	err = jpeg.Encode(b, resizeImage, nil)
	if err != nil {
		//log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		//log.Println(err)
		os.Exit(1)
	}
	//fmt.Print(http.FileServer(http.Dir("./public")))
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAICmLY9c6bDHAX", "qUbMQOIx0qRRV11zSRHKR7E2rjuIUP")
	if err != nil {
		panic(err)
		return

	}

	bucket, err := client.Bucket("lfo")
	if err != nil {
		panic(err)
		return

	}

	err = bucket.PutObjectFromFile(fileName, filePath)
	if err != nil {
		panic(err)
		return
	}

	var shopInfo model.ShopInfo
	shopInfo.ShopId = shopId
	shopInfo.Avatar = fileName
	shopInfo.UpdateTime = time.Now()
	orm.Update(&shopInfo)
	j, _ := json.Marshal(shopInfo)
	fmt.Fprint(w, string(j))
	return
	fmt.Fprint(w, j)
}

type RulerHasShop struct {
	model.ShopInfo `xorm:"extends"`
}

func FromRulerHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.FromUserAuth(r)
	if err != nil {
		panic(err)
		return
	}
	fmt.Print(user.UserId)
	results := make([]RulerHasShop, 0)
	orm.Table("shop_ruler").
		Desc("shop_info.create_time").
		Join("INNER", "shop_info", "shop_info.shop_id = shop_ruler.shop_id").
		Where("user_id = ?", user.UserId).
		Find(&results)
	if err != nil {
		panic(err)
	}
	j, _ := json.Marshal(results)
	fmt.Fprint(w, string(j))

}
