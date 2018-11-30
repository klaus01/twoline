package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	_ "github.com/klaus01/twoline/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

var gSessionID string

func GetRequest(url string) (*http.Request, error) {
	r, err := http.NewRequest("GET", url, nil)
	if len(gSessionID) > 0 {
		r.Header.Add("Sessionid", gSessionID)
	}
	return r, err
}

func PostRequest(url string, values *url.Values) (*http.Request, error) {
	bodyString := values.Encode()
	r, err := http.NewRequest("POST", url, bytes.NewBufferString(bodyString))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(bodyString)))
	if len(gSessionID) > 0 {
		r.Header.Add("Sessionid", gSessionID)
	}
	return r, err
}

// TestNewUserLogin 新用户登录
func TestNewUserLogin(t *testing.T) {
	data := url.Values{}
	data.Set("phonenumber", "18181994671")
	r, _ := PostRequest("/v1/users/login", &data)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	resString := w.Body.String()
	var resJSON map[string]interface{}
	Convey("新用户登录", t, func() {
		Convey("Code 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(len(resString), ShouldBeGreaterThan, 0)
		})
		Convey("返回了 JSON", func() {
			err := json.Unmarshal(w.Body.Bytes(), &resJSON)
			So(err, ShouldBeNil)
			So(resJSON, ShouldNotBeNil)
		})
		Convey("有 session", func() {
			data := resJSON["data"].(map[string]interface{})
			sessionID := data["sessionID"].(string)
			So(len(sessionID), ShouldBeGreaterThan, 0)
		})
		Convey("是新用户", func() {
			data := resJSON["data"].(map[string]interface{})
			isNewUser := data["isNewUser"].(bool)
			So(isNewUser, ShouldBeTrue)
		})
	})
}

// TestOldUserLogin 老用户登录
func TestOldUserLogin(t *testing.T) {
	data := url.Values{}
	data.Set("phonenumber", "18181994671")
	r, _ := PostRequest("/v1/users/login", &data)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	resString := w.Body.String()
	var resJSON map[string]interface{}
	Convey("老用户登录", t, func() {
		Convey("Code 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(len(resString), ShouldBeGreaterThan, 0)
		})
		Convey("返回了 JSON", func() {
			err := json.Unmarshal(w.Body.Bytes(), &resJSON)
			So(err, ShouldBeNil)
			So(resJSON, ShouldNotBeNil)
		})
		Convey("有 session", func() {
			data := resJSON["data"].(map[string]interface{})
			sessionID := data["sessionID"].(string)
			So(len(sessionID), ShouldBeGreaterThan, 0)
			gSessionID = sessionID
		})
		Convey("是老用户", func() {
			data := resJSON["data"].(map[string]interface{})
			isNewUser := data["isNewUser"].(bool)
			So(isNewUser, ShouldBeFalse)
		})
	})
}

// TestGetProfile 获取用户信息
func TestGetProfile(t *testing.T) {
	userID := 1

	r, _ := GetRequest("/v1/users/" + strconv.Itoa(userID))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	resString := w.Body.String()
	var resJSON map[string]interface{}
	Convey("获取用户信息", t, func() {
		Convey("Code 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(len(resString), ShouldBeGreaterThan, 0)
		})
		Convey("返回了 JSON", func() {
			err := json.Unmarshal(w.Body.Bytes(), &resJSON)
			So(err, ShouldBeNil)
			So(resJSON, ShouldNotBeNil)
		})
		Convey("ID = "+strconv.Itoa(userID), func() {
			data := resJSON["data"].(map[string]interface{})
			id := data["id"].(float64)
			So(id, ShouldEqual, userID)
		})
	})

	r, _ = GetRequest("/v1/users/333")
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	Convey("获取不正在的用户", t, func() {
		Convey("Code 404", func() {
			So(w.Code, ShouldEqual, 404)
		})
	})
}

// TestUserLogout 用户退出登录
func TestUserLogout(t *testing.T) {
	r, _ := GetRequest("/v1/users/logout")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	resString := w.Body.String()
	Convey("用户正常退出登录", t, func() {
		Convey("Code 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(len(resString), ShouldBeGreaterThan, 0)
		})
	})

	gSessionID = "xxxx"
	r, _ = GetRequest("/v1/users/logout")
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	resString = w.Body.String()
	Convey("不存在的 session 退出登录", t, func() {
		Convey("Code 401", func() {
			So(w.Code, ShouldEqual, 401)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(len(resString), ShouldBeGreaterThan, 0)
		})
	})

	gSessionID = ""
	r, _ = GetRequest("/v1/users/logout")
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	resString = w.Body.String()
	Convey("没 session 退出登录", t, func() {
		Convey("Code 401", func() {
			So(w.Code, ShouldEqual, 401)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(len(resString), ShouldBeGreaterThan, 0)
		})
	})
}
