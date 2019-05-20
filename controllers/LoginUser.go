package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"web_beego_dome01/controllers/database"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1022:07
 */

var (
	usernames string
	userpasswords string
)
type LoginUser struct {
	beego.Controller
}
func (this *LoginUser) Get(){

	this.DestroySession()		//销毁全部session
	this.Data["Github"]=Github
	this.Data["Email"]=Email
	this.TplName="login.html"
}
func (this *LoginUser) Post() {
	this.Data["Github"]=Github
	this.Data["Email"]=Email

	username:=this.GetString("Username")
	password:=this.GetString("Password")

	usernames=username
	userpasswords=password

	//set cookie
	this.Ctx.SetCookie("Username",username,100,"/")
	//this.Ctx.SetCookie("Password",password,100,"/")

	//session
	this.SetSession("Username",username)
	this.SetSession("Password",password)

	//登录用户判断

	if password== database.GetRedis(username){
		this.Redirect("/index",302)		//重定向
		return
	}
	this.DelSession("Username")		//销毁Username session
	this.Data["ErrorData"]="账号或密码错误，请重试"
	this.TplName="login.html"


}








