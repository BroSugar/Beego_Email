package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web_beego_dome01/controllers/database"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/2020:38
 */

type Deleter_User struct {
	beego.Controller
}

func (this *Deleter_User) Get()  {
	if this.GetSession("Username")==nil{
		this.Redirect("/index",302)
	}
	uname := this.GetSession("Username")
	upwd:=this.GetSession("Password")

	this.Data["username"] = uname
	this.Data["password"]=upwd
	uemail:=usernames+"email"
	this.Data["email"]=database.GetRedis(uemail)
	sex:=usernames+userpasswords
	this.Data["sex"]=database.GetRedis(sex)
	this.TplName="user_data_delete.html"
}

func (this *Deleter_User) Post(){
	uname:=this.Ctx.GetCookie("Username")
	fmt.Println(uname)
	if database.DelRedis(uname){
		if database.DeleteDB(uname){
			this.DelSession("Username")
			this.Redirect("/login",302)
		}
	}

	this.Redirect("/index",302)
}

