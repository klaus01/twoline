// @APIVersion 1.0.0
// @Title TwoLine API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact apple@amonster.net
// @TermsOfServiceUrl http://www.amonster.net/
package routers

import (
	"github.com/klaus01/twoline/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
