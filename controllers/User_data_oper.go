package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web_beego_dome01/controllers/database"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1721:22
 */
type User_data_oper struct {
	beego.Controller
}

func (this *User_data_oper) Get(){
	uname := this.GetSession("Username")
	upwd:=this.GetSession("Password")

	this.Data["username"] = uname
	this.Data["password"]=upwd
	uemail:=usernames+"email"
	this.Data["email"]=database.GetRedis(uemail)
	sex:=usernames+userpasswords
	this.Data["sex"]=database.GetRedis(sex)
	if uname == nil {
		this.Redirect("/login", 302)
		return
	}
	this.TplName="user_data_oper.html"
}

func (this *User_data_oper) Post(){
	uname:=this.Ctx.GetCookie("Username")
	upwd:=this.GetString("dpwd")
	uemail:=this.GetString("demail")
	usex:=this.GetString("dsex")
	fmt.Println("用户名:"+uname+" password:"+upwd+" email:"+uemail+" sex:"+usex)


	//修改用户信息
	if !database.Update(uname,upwd,uemail,usex){
		fmt.Println("更新失败")
		this.Redirect("/user_data_oper",302)
		return
	}
	user_oper_data:="操作成功"
	//this.SetSession("user_oper_data",user_oper_data)

	name := this.GetSession("Username")
	pwd:=this.GetSession("Password")

	this.Data["username"] = name
	this.Data["password"]=pwd
	this.Data["user_oper_data"]=user_oper_data
	this.Redirect("/user_data_oper",302)
}