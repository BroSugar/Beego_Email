package database

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1421:01
 */
var (
	dbhostsip  = "cdb-2dpvss9t.gz.tencentcdb.com"
	dbusername = "laohei"
	dbpassowrd = "Laohei123"
	dbname     = "web_user_data"
	port="10105"
	DB_Driver=dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+":"+port+")/"+dbname
)
type userlogin struct {
	uname string
	pps string
}
func init(){
	opend, db := OpenDB()
	if opend {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}
	QueryAlluser(db)
	db.Close()
}

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	CheckErr(err)
	return isOpen, db
}


func QueryAlluser(db *sql.DB){
	rows, err := db.Query("SELECT * FROM user")
	CheckErr(err)
	if err != nil {
		fmt.Println("error:", err)
	}
	for rows.Next() {
		var uid int
		var uemail string
		var uname string
		var upwd string
		var usex string
		CheckErr(err)
		err = rows.Scan(&uid, &uemail, &uname, &upwd, &usex)
		sexdata:=uname+upwd		//性别
		if !SetRedis(uname,upwd){
			fmt.Println("导出到redis失败")
			return
		}
		if !SetRedis(sexdata,usex) {
			fmt.Println("导出到redis失败")
			return
		}
		useremail:=uname+"email"
		if !SetRedis(useremail,uemail){
			fmt.Println("导出到redis失败")
			return
		}
	}
}

//条件查询
func Query(username string) string{
	conn,err := sql.Open("mysql",dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+":"+port+")/"+dbname)

	if err!=nil{
		beego.Info("连接失败")
	}
	defer conn.Close()

	query:="select * from User where Username='"+username+"'"
	row ,err:= conn.Query(query)
	if err!=nil{
		panic(err)
	}
	var id int
	var email string
	var uname string
	var upassword string
	var sex string
	for row.Next(){
		row.Columns()
		err=row.Scan(&id,&email,&uname,&upassword,&sex)
		if err!=nil{
			panic(err)
		}
	}
	return uname
}

func InsertDB(upwd string,uemail string,uname string,usex string) bool{
	opend, db := OpenDB()
	if opend {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}
	if len(Query(uname))!=0{
		fmt.Println("用户存在")
		return false
	}
	if !insertToDB(db,uname,uemail,upwd,usex){
		fmt.Println("插入失败")
		return false
	}
	db.Close()
	return true
}

//增加
func insertToDB(db *sql.DB,uname string,uemail string,upwd string,usex string) bool{
	stmt, err := db.Prepare("insert user set Username=?,Email=?,Password=?,Sex=?")
	CheckErr(err)
	res, err := stmt.Exec(uname,uemail, upwd,usex)
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	if err != nil {
		fmt.Println("插入数据失败")
		return false
	}
		fmt.Println("插入数据成功：", id)
	return true
}

func DeleteDB(uname string) bool{
	opend, db := OpenDB()
	if opend {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}

	//删除
	if !DeleteFromDB(db,uname){
		fmt.Println("删除操作失败")
		db.Close()
		return false
	}
	db.Close()
	return true
}

//删除
func DeleteFromDB(db *sql.DB, autid string) bool{
	stmt, err := db.Prepare("delete from user where Username=?")
	CheckErr(err)
	res, err := stmt.Exec(autid)
	affect, err := res.RowsAffected()
	fmt.Println("删除数据：", affect)
	if affect<1 {
		return false
	}
	return true
}

func Update (uname string,upwd string,uemail string,usex string) bool {
	opend, db := OpenDB()
	if opend {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}
	if !UpdateDB(db,uname,upwd,uemail,usex){
		return false
	}
	db.Close()
	return true
}

//更新

//UpdateDB
func UpdateDB (db *sql.DB, uname string,upwd string,uemail string,usex string) bool{
	stmt, err := db.Prepare("update user set Email=?,Password=?,Sex=? where Username=?")
	CheckErr(err)
	res, err := stmt.Exec(uemail,upwd,usex, uname)
	affect, err := res.RowsAffected()
	fmt.Println("更新数据：", affect)
	if affect<1 {
		fmt.Println("更新失败")
		return false
	}else {
		val:=uname+upwd
		if !SetRedis(val,usex){
			fmt.Println("导出到redis失败")
			return false
		}
	}
	CheckErr(err)
	return true
}








//工具类
func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	fmt.Println(t)
	return str
}

func GetMD5Hash(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}

func GetNowtimeMD5() string {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	return GetMD5Hash(timestamp)
}