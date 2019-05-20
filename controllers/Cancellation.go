package controllers
//注销
import (
	"github.com/astaxie/beego"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1022:30
 */

type Cancellation struct {
	beego.Controller
}

func (this *Cancellation) Get() {
	this.DelSession("Username")
	this.Redirect("/login",302)
}