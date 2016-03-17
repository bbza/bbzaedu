package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//
	this.Redirect("home/index", 301)
	this.TplName = "index.html"
}
