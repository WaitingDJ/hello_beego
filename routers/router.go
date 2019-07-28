package routers

import (
	"github.com/astaxie/beego"
	"hello_beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{}, "get:GetRegister;post:PostRegister")
	beego.Router("/login", &controllers.MainController{}, "get:GetLogin;post:PostLogin")
}
