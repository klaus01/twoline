package controllers

import (
	"net/http"
)

type ErrorController struct {
	AppController
}

func (c *ErrorController) Error404() {
	c.ResultErrorAndMessage(http.StatusNotFound, "请求资源不存在")
}

func (c *ErrorController) ErrorDb() {
	c.ResultErrorAndMessage(http.StatusInternalServerError, "数据库错误")
}
