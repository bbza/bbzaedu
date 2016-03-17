package controllers

//这是导入的包
import (
	//"fmt"
	"github.com/astaxie/beego"
)

//zheshi kongzhiqi tou
type AdminController struct {
	beego.Controller
}

//admin/index
func (this *AdminController) Index() {
	//
	this.Data["title"] = "这是后台首页"     //这里是模板变量输出
	this.TplName = "admin/index.html" //templater url
}

//admin/login.html

func (this *AdminController) Login() {
	//if this.Ctx.Input.Method() == "POST" {
	//Zz这里进行验证
	//fmt.Println(this.Input())
	// this.Redirect("index", 302) //重定向
	//}
	//这里判断post
	this.TplName = "admin/login.html"
}
