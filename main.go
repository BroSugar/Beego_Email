package main

import (
	"container/list"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
	"web_beego_dome01/controllers"
	_ "web_beego_dome01/routers"
)

var (
	dbhostsip  = "localhsot"
	dbusername = "root"
	dbpassowrd = "Xiaoye@0310"
	dbname     = "web_user_data"
	port="3306"
)
func query() {
	db, err := sql.Open("mysql", "root:Xiaoye@0310@/web_user_data?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	//    //普通demo
	names:=list.New()
	pass:=list.New()
	for rows.Next() {
		var userid int
		var useremail string
		var username string
		var userpassword string
		var usersex string

		rows.Columns()
		err = rows.Scan(&userid,&useremail, &username, &userpassword, &usersex)
		checkErr(err)
		//setRedis(usernmae,userpassword)
		names.PushBack(username)
		pass.PushBack(userpassword)
		//fmt.Println(userid)
		//fmt.Println(useremail)
		//fmt.Println(username)
		//fmt.Println(userpassword)
		//fmt.Println(usersex)
	}
	fmt.Println("数组长度为：")
	fmt.Println(names.Len())
	for i:=0;i<names.Len();i++{

	}
	fmt.Println("账号")
	for e:=names.Front();e!=nil;e=e.Next()  {
		fmt.Println(e.Value)
	}
	fmt.Println("密码")
	for e:=pass.Front();e!=nil ;e=e.Next()  {
		fmt.Println(e.Value)
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
func main() {
	//
	beego.ErrorController(&controllers.ErrorController{})		//错误页面
	beego.Run()

}

