package routers

import (
	"beego-bbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{}, "*:Index")
	beego.AutoRouter(&controllers.PostController{})
}
