package models

import (
	"container/list"
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1421:01
 */
var (
	dbhostsip  = "localhsot"
	dbusername = "laohei"
	dbpassowrd = "Laohei123"
	dbname     = "web_user_data"
	port="3306"
)
type userlogin struct {
	uname string
	pps string
}
func queryAlluser(){

	conn,err := sql.Open("mysql",dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+":"+port+")/"+dbname)

	if err!=nil{
		beego.Info("连接失败")
	}
	defer conn.Close()

	query:="select * from user"
	row ,err:= conn.Query(query)
	if err!=nil{
		panic(err)
	}
	var id int
	var email string
	var uname string
	var upassword string
	var sex string

	names:=list.New()
	pass:=list.New()
	for row.Next(){
		row.Columns()
		err=row.Scan(&id,&email,&uname,&upassword,&sex)
		names.PushBack(uname)
		pass.PushBack(upassword)
		if err!=nil{
			panic(err)
		}
	}
}
//条件查询
func query(username string){

	conn,err := sql.Open("mysql",dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+":"+port+")/"+dbname)

	if err!=nil{
		beego.Info("连接失败")
	}
	defer conn.Close()

	query:="select * from User where Username="+username
	row ,err:= conn.Query(query)
	if err!=nil{
		panic(err)
	}
	var id int
	var email string
	var uname string
	var upassword string
	var sex string
	var u userlogin
	for row.Next(){
		row.Columns()
		err=row.Scan(&id,&email,&uname,&upassword,&sex)
		u.uname=uname
		u.pps=upassword
		if err!=nil{
			panic(err)
		}
	}
}
//增加
func add(){

}
//删除
func delete(){

}
//更新
func update(){

}