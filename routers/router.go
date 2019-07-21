package routers

import (
	"hello_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/abc", &controllers.MainController{})
}
