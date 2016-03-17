package controllers

import (
	"github.com/astaxie/beego"
)

type TeacherController struct {
	beego.Controller
}

func (this *TeacherController) Index() {
	this.Data["list"] = "教师列表"
}

//teacher/show.html?tid=1
/*
 *直接获取教师详情
 * tid  //教师id
 */
func (this *TeacherController) Show() {
	if this.GetString("tid") != nil {
		var tid = this.GetString("tid")
		//这里根据model获取教师详情
		//switch
	}
	this.TplName = "teacher/show.html"
}
