package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id   int
	Name string
}

func (u *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}
