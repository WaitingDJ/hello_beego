package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"hello_beego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetRegister() {
	c.TplName = "register.html"

	// 插入数据
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Name = "testName"
	//user.Pwd = "123456"
	//_,err := o.Insert(&user)
	//if err != nil {
	//	logs.Info("插入失败", err)
	//	return
	//}

	// 查询数据
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 2
	//err := o.Read(&user)
	//if err != nil {
	//	logs.Info("查询失败", err)
	//	return
	//}
	//logs.Info("查询成功", user)

	// 更新数据
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 2
	//err := o.Read(&user)
	//if err == nil {
	//	user.Name = "testName2"
	//	user.Pwd = "234567"
	//	_,err := o.Update(&user)
	//	if err != nil {
	//		logs.Info("更新失败", err)
	//	}
	//	logs.Info("更新成功")
	//}

	// 删除数据
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 2
	//
	//err := o.Read(&user)
	//if err != nil {
	//	logs.Info("没有查找到当前项")
	//} else {
	//	logs.Info("查找完毕")
	//	_, err := o.Delete(&user)
	//	if err != nil {
	//		logs.Info("删除错误", err)
	//		return
	//	}
	//	logs.Info("删除成功")
	//}
	//
	//c.Data["msg"] = "beego.me"
	//c.TplName = "index.html"
}

func (c *MainController) PostRegister() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	logs.Info(userName, pwd)
	fmt.Println(81)
	if userName == "" || pwd == "" {
		fmt.Println(83)
		logs.Info("数据不能为空")
		c.Redirect("/register", 302)
		return
	}

	o := orm.NewOrm()
	user := models.User{Name: userName, Pwd: pwd}
	//user.Name = userName
	//user.Pwd = pwd
	_, err := o.Insert(&user)
	//err := o.Read(&user, "name")
	//fmt.Println(94)
	if err != nil {
		c.Ctx.WriteString("注册失败")
		return
		//_, err = o.Insert(&user)
		//if err != nil {
		//	c.Ctx.WriteString("注册失败")
		//	return
		//}
		//c.Ctx.WriteString("注册成功")
		//c.Redirect("/login", 200)
		//return
	}
	c.Ctx.WriteString("注册成功")

	//c.Data["msg"] = "我变啦,怎么滴~"
	//c.TplName = "index.html"
}

func (c *MainController) GetLogin() {
	c.TplName = "login.html"
}
func (c *MainController) PostLogin() {
	userName := c.GetString("userName")
	//pwd := c.GetString("pwd")

	o := orm.NewOrm()
	user := models.User{}
	//user.Pwd = pwd
	user.Name = userName

	err := o.Read(&user, "Name")

	if err != nil {
		c.Ctx.WriteString("无此用户")
	}
	c.Ctx.WriteString("登录成功")
}
