package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Get("/api", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Router("/json", &controllers.ApiController{})

	beego.Router("/xml", &controllers.ApiController{}, "get:GetPerson")
}
