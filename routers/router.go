package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//beego.Router("/", &controllers.MainController{})
	//
	//
	//
	//beego.Get("/demo" , func(context *context.Context) {
	//	context.OutPut.Body([] byte("demo test"))
	//})

	beego.Router("/api/list", &controllers.MainController{}, "*:ListFood")

	beego.Router("/api/posttest", &controllers.MainController{}, "*:TestPost")
	//beego.Router("/api/posttest", &controllers.MainController{}, "Post:TestPost")
}
