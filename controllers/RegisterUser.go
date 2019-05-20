package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web_beego_dome01/controllers/database"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1022:08
 */

type RegisterUser struct {
	beego.Controller
}

func (this *RegisterUser) Get(){
	register_err:=this.GetSession("register")
	this.Data["register_err"]=register_err
	this.TplName="register.html"

}

func (this *RegisterUser) Post(){

	uname:=this.GetString("uname")
	upassword:=this.GetString("upassword")
	uemail:=this.GetString("uemail")
	usex:=this.GetString("usex")

	fmt.Println(uname,uemail,upassword,usex)
	if usex=="男" || usex=="女" {
		if len(uname)==0 && len(upassword)==0 && len(uemail)==0 && len(usex)==0 && uname=="" && upassword=="" && uemail=="" && usex==""{
			register_err:="未获取到用户数据。"
			this.SetSession("register",register_err)
			this.Redirect("/register",302)
			return
		}else {
			if !database.InsertDB(upassword,uemail,uname,usex){
				register_err:="用户存在请重试。"
				this.SetSession("register",register_err)
				this.Redirect("/register",302)
				return
			}
		}
		this.DelSession("register")
		this.Redirect("/login",302)
	}
	this.Redirect("/register",302)


}