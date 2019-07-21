package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"hello_beego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	o := orm.NewOrm()

	user := models.User{}

	user.Name = "testName"
	user.Pwd = "123456"

	_,err := o.Insert(&user)

	if err != nil {
		logs.Info("插入失败", err)
		return
	}

	c.Data["msg"] = "beego.me"
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	c.Data["msg"] = "我变啦,怎么滴~"
	c.TplName = "index.html"
}
