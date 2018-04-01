package user

import (
	"cafe.lsfoo.com/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:lsf000000@/cxj?charset=utf8")
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

type WXJscode2Session struct {
	SessionKey string `json:"session_key"`
	ExpiresIn  int    `json:"expires_in"`
	OpenId     string `json:"openid"`
}

//通过验证获取用户资源
func FromAuthorization(r *http.Request) (*model.User, error) {
	//获取微信code
	code, err := FromAuthHeader(r)
	if err != nil {
		panic(err)
	}
	//获取openid
	wx, err := FromWXCode(code)
	if err != nil {
		panic(err)
	}
	openid := wx.OpenId
	//通过openid获取用户资源

	user := &model.User{WxOpenId: openid, CreateTime: time.Now()}
	has, err := engine.Get(user)
	if !has {
		engine.Insert(user)
	}
	//用户存在返回用户资源

	return user, err

}

func FromWXCode(code string) (WXJscode2Session, error) {

	var r WXJscode2Session
	uri := "https://api.weixin.qq.com/sns/jscode2session?appid=wx32301b896c4287d0&secret=b44d359c84bf958f327c38873b091795&js_code=" + code + "&grant_type=authorization_code"
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
