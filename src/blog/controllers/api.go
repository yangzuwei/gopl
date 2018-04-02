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
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
}

type ApiController struct {
	beego.Controller
}

type person struct {
	Name   string
	Age    int
	Height int
}

func (this *ApiController) Get() {
	o := orm.NewOrm()
	user := models.Users{Id: 2}
	err := o.Read(&user)
	fmt.Println(err)
	fmt.Println(user)
	this.Data["json"] = &user
	this.ServeJSON()
}

func (this *ApiController) GetPerson() {
	xiaoming := &person{"xiaoming", 12, 150}
	this.Data["json"] = xiaoming
	this.ServeJSON()
}
