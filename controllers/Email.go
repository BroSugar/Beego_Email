package controllers

import (
	"github.com/astaxie/beego"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1720:51
 */

type Email_Oper struct {
	beego.Controller
}



func (this *Email_Oper) Get(){
	uname := this.GetSession("Username")
	this.Data["username"] = uname




	this.TplName="mail_compose.html"
}
