package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

//这就是首页.对于你我已经无语了.
func (this *HomeController) Index() {
	this.Data["title"] = "**首页"
	this.TplName = "home/index.html"
}
