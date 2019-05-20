package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1922:08
 */

type Write_Email struct {
	beego.Controller
}

func (this *Write_Email) Get(){

	val:=this.GetString("write_email")
	fmt.Println(val)

	this.Redirect("/mail_compose",302)
}