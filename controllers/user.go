package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/klaus01/twoline/services"
	"github.com/klaus01/twoline/until"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// // @Title CreateUser
// // @Description create users
// // @Param	body		body 	models.User	true		"body for user content"
// // @Success 200 {int} models.User.Id
// // @Failure 403 body is empty
// // @router / [post]
// func (u *UserController) Post() {
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
// func (u *UserController) GetAll() {
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
// func (u *UserController) Get() {
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
// func (u *UserController) Put() {
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
// func (u *UserController) Delete() {
// 	uid := u.GetString(":uid")
// 	models.DeleteUser(uid)
// 	u.Data["json"] = "delete success!"
// 	u.ServeJSON()
// }

// Login 登录
// @Title Login
// @Description Logs user into the system
// @Param	phonenumber		query 	string	true		"手机号"
// @Success 200 {object} models.User
// @Failure 400 缺少参数
// @router /login [get]
func (u *UserController) Login() {
	phonenumber := strings.TrimSpace(u.GetString("phonenumber"))
	if len(phonenumber) <= 0 {
		until.ResultError(&u.Controller, 400, "缺少参数")
		return
	}
	user, err := services.Login(phonenumber)
	if err != nil {
		until.ResultError(&u.Controller, 500, err.Error())
		return
	}
	if user == nil {
		until.ResultError(&u.Controller, 500, "内部错误")
		return
	}
	until.ResultData(&u.Controller, "登录成功", &user)
}

// // @Title logout
// // @Description Logs out current logged in user session
// // @Success 200 {string} logout success
// // @router /logout [get]
// func (u *UserController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }
