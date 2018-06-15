package routers

import (

	"loginServer/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/twelve",

		beego.NSNamespace("/user",

			beego.NSInclude(

				&controllers.UserController{},

			),

		),

	)

	beego.AddNamespace(ns)

}
