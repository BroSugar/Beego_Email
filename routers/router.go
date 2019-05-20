package routers

import (
	"github.com/astaxie/beego"
	"web_beego_dome01/controllers"
)

//过滤器


func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.Index{})

	beego.Router("/login",&controllers.LoginUser{})	//登录
	beego.Router("/register",&controllers.RegisterUser{})		//注册
	beego.Router("/cancellation",&controllers.Cancellation{})		//注销

	beego.Router("/mail_compose",&controllers.Email_Oper{})		//写邮件
	beego.Router("/user_data_oper",&controllers.User_data_oper{})		//用户个人信息修改

	beego.Router("/set_email",&controllers.Write_Email{})			//邮件

	beego.Router("/user_oper_delete",&controllers.Deleter_User{})		//删除账号
}
