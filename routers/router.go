package routers

import (
	"bbzaedu/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("", &controllers.MainController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.ShowcurriculumController{})
}
