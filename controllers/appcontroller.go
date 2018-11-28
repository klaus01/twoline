package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

const (
	SessionKeyUserID = "userID"
)

type AppController struct {
	beego.Controller
}

type resultData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *AppController) ResultData(message string, data interface{}) {
	c.Data["json"] = &resultData{http.StatusOK, message, data}
	c.ServeJSON()
}

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

func (c *AppController) ResultErrorAndMessage(code int, message string) {
	c.Ctx.Output.SetStatus(code)
	c.Data["json"] = &resultData{code, message, nil}
	c.ServeJSON()
}

func (c *AppController) SetSessionUserID(id int64) {
	c.SetSession(SessionKeyUserID, id)
}

func (c *AppController) GetSessionUserID() int64 {
	if userID := c.GetSession(SessionKeyUserID); userID != nil {
		return userID.(int64)
	}
	c.ResultError(http.StatusUnauthorized)
	c.StopRun()
	return 0
}
