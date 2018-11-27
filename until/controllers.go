package until

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

func ResultData(c *beego.Controller, message string, data interface{}) {
	c.Data["json"] = &resultData{200, message, data}
	c.ServeJSON()
}

func ResultError(c *beego.Controller, code int, message string) {
	c.Ctx.Output.SetStatus(code)
	c.Data["json"] = &resultData{code, message, nil}
	c.ServeJSON()
}

type resultData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
