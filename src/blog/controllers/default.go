package controllers

import (
	"blog/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	user := models.Users{Id: 2}
	err := o.Read(&user)
	fmt.Println(err)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = user.Name
	c.TplName = "index.tpl"
}
