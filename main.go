package main

import (
	_ "hello_beego/routers"
	_ "hello_beego/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

