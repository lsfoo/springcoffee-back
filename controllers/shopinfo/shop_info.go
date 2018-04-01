package shopinfo

import (
	//"cafe.lsfoo.com/controllers/shop"
	"cafe.lsfoo.com/auth"
	"cafe.lsfoo.com/db"
	"cafe.lsfoo.com/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var orm = db.Orm
var key = "fukkkkyou"

func FromRulerQrhandler(w http.ResponseWriter, r *http.Request) {
	//var png []byte
	shop, err := auth.FromShopAuth(r)
	if err != nil {
		return
	}
	fmt.Print(shop)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "jwt"
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Second * 60).Unix() //token超时时间
	claims["sid"] = shop.ShopId
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	png, _ := qrcode.Encode(tokenString, qrcode.Medium, 256)
	//fmt.Fprint(w, tokenString)
	fmt.Fprint(w, string(png))
}

type QRresult struct {
	Result string
}

func NewRulerHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.FromUserAuth(r)
	if err != nil {
		panic(err)
		return
	}
	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)

	var qrResult QRresult
	json.Unmarshal([]byte(body), &qrResult)
	//fmt.Print(string(body))
	token := qrResult.Result
	fmt.Print(token)
	claims, err := AuthRulerQr(token)
	if err != nil {
		//	panic(err)
		return
	}
	var shopId int
	shopId = int(claims["sid"].(float64))
	var ruler model.ShopRuler
	ruler.UserId = user.UserId
	ruler.ShopId = shopId
	ruler.CreateTime = time.Now()
	orm.Insert(&ruler)
	fmt.Fprint(w, ruler)

}
func AuthRulerQr(qr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(qr, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("Timing is everything")
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func FromGpsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	var si []model.ShopInfo
	orm.Find(&si)
	jsons, _ := json.Marshal(si)
	fmt.Fprint(w, string(jsons))
}
func FromIdHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	shopInfo := model.ShopInfo{ShopId: id}
	orm.Get(&shopInfo)
	jsons, _ := json.Marshal(shopInfo)
	fmt.Fprint(w, string(jsons))
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	shop, _ := auth.FromShopAuth(r)
	if shop.ShopId == 0 {
		fmt.Fprint(w, 0)
		return
	}
	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	var shopInfo model.ShopInfo
	json.Unmarshal(body, &shopInfo)
	orm.Insert(&shopInfo)
	fmt.Fprint(w, shopInfo)
}
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	shop, err := auth.FromShopAuth(r)
	if err != nil {
		panic(err)
		return
	}
	resp := r.Body
	r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	var shopInfo model.ShopInfo
	shopInfo.ShopId = shop.ShopId
	json.Unmarshal(body, &shopInfo)
	jsons := db.Update(shopInfo)
	fmt.Fprint(w, jsons)
}
func UpdatePositionHandler(w http.ResponseWriter, r *http.Request) {
	shop, err := auth.FromShopAuth(r)
	if err != nil {
		return
	}
	resp := r.Body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(resp)
	fmt.Println(string(body))
	var shopInfo model.ShopInfo
	json.Unmarshal([]byte(body), &shopInfo)
	fmt.Print(shopInfo)
	shopInfo.ShopId = shop.ShopId
	shopInfo.UpdateTime = time.Now()
	//fmt.Print(shopInfo)
	jsons := db.Update(&shopInfo)
	fmt.Fprint(w, jsons)
}
