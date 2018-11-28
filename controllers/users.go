package controllers

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/klaus01/twoline/models"
	"github.com/klaus01/twoline/services"
)

// Operations about Users
type UsersController struct {
	AppController
}

type LoginResponseJSON struct {
	SessionID string `json:"sessionID"`
	IsNewUser bool   `json:"isNewUser"`
}

// Login 登录
// @Title Login
// @Description 登录或注册
// @Param	phonenumber		query 	string	true		"手机号"
// @Success 200 {object} controllers.LoginResponseJSON
// @Failure 400 缺少参数
// @router /login [post]
func (u *UsersController) Login() {
	phonenumber := strings.TrimSpace(u.GetString("phonenumber"))
	if len(phonenumber) <= 0 {
		u.ResultError(http.StatusBadRequest)
		return
	}
	user, isNewUser, err := services.Login(phonenumber)
	if err != nil {
		beego.Error(err)
		u.ResultError(http.StatusInternalServerError)
		return
	}
	if user == nil || user.ID <= 0 {
		beego.Error("登录注册时 user 为空", user)
		u.ResultError(http.StatusInternalServerError)
		return
	}
	u.SetSessionUserID(user.ID)
	data := &LoginResponseJSON{u.CruSession.SessionID(), isNewUser}
	u.ResultData("登录成功", data)
}

// Logout 退出登录
// @Title logout
// @Description 退出登录，销毁 sessionID
// @Success 200 {nil} logout success
// @Failure 401 未登录或登录已过期
// @router /logout [get]
func (u *UsersController) Logout() {
	defer u.DestroySession()
	u.GetSessionUserID()
	u.ResultData("登出成功", nil)
}

// Profile 获取用户信息
// @Title Get
// @Description 获取用户信息
// @Param	uid		path 	string	true		"用户ID"
// @Success 200 {object} models.User
// @Failure 400 :uid 为空
// @Failure 401 未登录或登录已过期
// @Failure 404 :uid 不存在
// @router /:uid [get]
func (u *UsersController) Profile() {
	uid, err0 := u.GetInt64(":uid")
	if err0 != nil {
		u.ResultErrorAndMessage(http.StatusBadRequest, err0.Error())
		return
	}
	loginUserID := u.GetSessionUserID()
	user, err := models.GetUserByID(uid, loginUserID)
	if err != nil {
		beego.Error(err)
		u.ResultError(http.StatusInternalServerError)
		return
	}
	if user == nil {
		u.ResultErrorAndMessage(http.StatusNotFound, "请求的用户ID不存在")
		return
	}
	u.ResultData("", user)
}

// // @Title CreateUser
// // @Description create users
// // @Param	body		body 	models.User	true		"body for user content"
// // @Success 200 {int} models.User.Id
// // @Failure 403 body is empty
// // @router / [post]
// func (u *UsersController) Post() {
// 	var user models.User
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 	uid := models.AddUser(user)
// 	u.Data["json"] = map[string]string{"uid": uid}
// 	u.ServeJSON()
// }

// // @Title GetAll
// // @Description get all Users
// // @Success 200 {object} models.User
// // @router / [get]
// func (u *UsersController) GetAll() {
// 	users := models.GetAllUsers()
// 	arrayUsers := make([]*models.User, len(users))
// 	for _, v := range users {
// 		arrayUsers = append(arrayUsers, v)
// 	}
// 	u.Data["json"] = map[string][]*models.User{"users": arrayUsers}
// 	u.ServeJSON()
// }

// // @Title Get
// // @Description get user by uid
// // @Param	uid		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.User
// // @Failure 403 :uid is empty
// // @router /:uid [get]
// func (u *UsersController) Get() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		user, err := models.GetUser(uid)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = user
// 		}
// 	}
// 	u.ServeJSON()
// }

// // @Title Update
// // @Description update the user
// // @Param	uid		path 	string	true		"The uid you want to update"
// // @Param	body		body 	models.User	true		"body for user content"
// // @Success 200 {object} models.User
// // @Failure 403 :uid is not int
// // @router /:uid [put]
// func (u *UsersController) Put() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		var user models.User
// 		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 		uu, err := models.UpdateUser(uid, &user)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = uu
// 		}
// 	}
// 	u.ServeJSON()
// }

// // @Title Delete
// // @Description delete the user
// // @Param	uid		path 	string	true		"The uid you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 uid is empty
// // @router /:uid [delete]
// func (u *UsersController) Delete() {
// 	uid := u.GetString(":uid")
// 	models.DeleteUser(uid)
// 	u.Data["json"] = "delete success!"
// 	u.ServeJSON()
// }
