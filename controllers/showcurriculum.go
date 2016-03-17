package controllers

import (
	"github.com/astaxie/beego"
)

type ShowcurriculumController struct {
	beego.Controller
}

//下面是请求方法 /user/home.html
func (this *ShowcurriculumController) Pupillist() {
	this.TplName = "showcurriculum/pupillist.html"
}

func (this *ShowcurriculumController) Juniorlist() {
	this.TplName = "showcurriculum/juniorlist.html"
}

func (this *ShowcurriculumController) Collegelist() {
	this.TplName = "showcurriculum/collegelist.html"
}
