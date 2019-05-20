package controllers

import (
	"github.com/astaxie/beego"
	"math/rand"
)
var Github="XuanQieXiaoYe"
var Email="xiaoye_0310@126.com"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Github"]=Github
	c.Data["Email"]=Email
	c.TplName = "landing.html"
}





type LIKE struct {
	Food string
	Watch string
	Listen string
}



type Index struct {
	beego.Controller
}
func (this *Index) Get() {
	this.Data["Github"]=Github
	this.Data["Email"]=Email
	uname := this.GetSession("Username")
	this.Data["username"] = uname
	if uname == nil {
		this.Redirect("/login", 302)
		return
	}


	this.Data["Total"] = rand.Int31n(100)
	this.Data["FreeNum"] = rand.Int31n(200)
	this.Data["UsedNum"] = rand.Int31n(200)

	this.TplName = "mail_compose.html"
}
