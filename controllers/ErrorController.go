package controllers

import (
	"github.com/astaxie/beego"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1023:44
 */
type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.TplName="404.html"
}
func (this *ErrorController) Error500()  {
	this.TplName="500.html"
}