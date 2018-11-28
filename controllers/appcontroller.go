package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

const (
	// SessionKeyUserID Session 中存 UserID 的 key
	SessionKeyUserID = "userID"
)

// AppController 封装应用 controller 层公共方法
type AppController struct {
	beego.Controller
}

type resultData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResultData 返回数据，http status=200
func (c *AppController) ResultData(message string, data interface{}) {
	c.Data["json"] = &resultData{http.StatusOK, message, data}
	c.ServeJSON()
}

// ResultError 返回错误码，http.Status 定义的常量
func (c *AppController) ResultError(code int) {
	message := "服务端未知错误"
	switch code {
	case http.StatusBadRequest:
		message = "缺少参数"
	case http.StatusUnauthorized:
		message = "未登录或登录已过期"
	case http.StatusInternalServerError:
		message = "服务器内部错误"
	}
	c.ResultErrorAndMessage(code, message)
}

// ResultErrorAndMessage 返回错误码和提示信息，http.Status 定义的常量
func (c *AppController) ResultErrorAndMessage(code int, message string) {
	c.Ctx.Output.SetStatus(code)
	c.Data["json"] = &resultData{code, message, nil}
	c.ServeJSON()
}

// SetSessionUserID 保存 UserID 到 session 中
func (c *AppController) SetSessionUserID(id int64) {
	c.SetSession(SessionKeyUserID, id)
}

// GetSessionUserID 从 session 中获取 UserID，如果未设置 session 或已过期，则会返回 401 并终止请求
func (c *AppController) GetSessionUserID() int64 {
	if userID := c.GetSession(SessionKeyUserID); userID != nil {
		return userID.(int64)
	}
	c.ResultError(http.StatusUnauthorized)
	c.StopRun()
	return 0
}
