package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//下面是请求方法 /user/home.html
func (this *UserController) Home() {
	this.TplName = "user/index.html"
}
