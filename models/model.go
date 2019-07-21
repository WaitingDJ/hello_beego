package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	)

type User struct {
	Id int
	Name string
	Pwd string
}

func init() {
	// 设置基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	// 映射Model数据
	orm.RegisterModel(new(User))
	// 生成表
	orm.RunSyncdb("default", false, true)
}