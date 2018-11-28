package controllers

import (
	"net/http"
)

// ErrorController 应用全局错误处理
type ErrorController struct {
	AppController
}

// Error404 404 错误处理
func (c *ErrorController) Error404() {
	c.ResultErrorAndMessage(http.StatusNotFound, "请求资源不存在")
}

// ErrorDb 数据库错误处理
func (c *ErrorController) ErrorDb() {
	c.ResultErrorAndMessage(http.StatusInternalServerError, "数据库错误")
}
