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

	//这里获取推荐的用户 根据model获取.并填充到首页
	//优质教师
	//
	//大学生
	//名师
	this.TplName = "home/index.html"
}
