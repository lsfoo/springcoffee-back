package auth

import (
	"cafe.lsfoo.com/db"
	"cafe.lsfoo.com/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	//	"time"
)

var orm = db.Orm

type WXJscode2Session struct {
	SessionKey string `json:"session_key"`
	ExpiresIn  int    `json:"expires_in"`
	OpenId     string `json:"openid"`
}

//通过验证获取用户资源
func FromShopAuth(r *http.Request) (*model.Shop, error) {

	openid := GetOpenId(r)

	//通过openid获取用户资源
	shop := &model.Shop{WxOpenId: openid}
	has, _ := orm.Get(shop)
	if !has {
		return shop, errors.New("不存在的商家")
	}
	return shop, nil
}

//通过验证获取用户资源
func FromUserAuth(r *http.Request) (*model.User, error) {
	//获取微信code

	openid := GetOpenId(r)

	//通过openid获取用户资源
	user := &model.User{WxOpenId: openid}
	has, err := orm.Get(user)
	if !has {
		user.CreateTime = time.Now()
		orm.Insert(user)
	}
	return user, err
}
func GetOpenId(r *http.Request) string {
	code, err := FromAuthHeader(r)
	if err != nil {
		panic(err)
	}
	//获取openid
	wx, err := FromWXCode(code)
	if err != nil {
		panic(err)
	}
	return wx.OpenId

}

func FromWXCode(code string) (WXJscode2Session, error) {

	var r WXJscode2Session
	//获取openid的uri
	uri := "https://api.weixin.qq.com/sns/jscode2session?appid=wx32301b896c4287d0&secret=b44d359c84bf958f327c38873b091795&grant_type=authorization_code&js_code=" + code
	resp, err := http.Get(uri)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &r)
	return r, err
}

//获取头部Authorization 方式 bearer

func FromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}
	return authHeaderParts[1], nil
}
