package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["loginServer/controllers:UserController"] = append(beego.GlobalControllerRouter["loginServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["loginServer/controllers:UserController"] = append(beego.GlobalControllerRouter["loginServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["loginServer/controllers:UserController"] = append(beego.GlobalControllerRouter["loginServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Regist",
			Router: `/regist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["loginServer/controllers:UserController"] = append(beego.GlobalControllerRouter["loginServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdatePwd",
			Router: `/updatepwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
