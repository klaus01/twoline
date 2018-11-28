package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/klaus01/twoline/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/klaus01/twoline/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Profile",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/klaus01/twoline/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/klaus01/twoline/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/klaus01/twoline/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/klaus01/twoline/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
