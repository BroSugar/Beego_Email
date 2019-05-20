package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
*@Author:XuanQieXiaoYe
*@Date:2019/5/1421:00
 */


var (
	typecon="tcp"
	Redisport="192.168.0.120:6379"
)



func multiRedis(){
	//事务开始
}
func execRedis(){
	//事务提交
}

func setRedis(key string,value string)  {
	//Redis数据库连接
	c,err:=redis.Dial(typecon,Redisport)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer c.Close()
	//通过DO函数发送，redis命令
	v,err:=c.Do("set",key,value)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
func getRedis(key string,value string){
	//读取列表
	//Redis数据库连接
	c,err:=redis.Dial(typecon,Redisport)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer c.Close()
	values,_:=redis.Values(c.Do("get",key,value))
	var s1 string
	redis.Scan(values,&s1)
	fmt.Println(s1)
}
func lpushRedis(key string,val string){
	////Redis数据库连接
	//c,err:=redis.Dial(typecon,Redisport)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//defer c.Close()
	////通过DO函数发送，redis命令
	//v,err:=c.Do("set",key,val)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(v)
	//
	////操作列表
	//c.Do("lpush",key,"qqq")
	//c.Do("lpush","redis","www")
	//c.Do("lpush","redis","eee")
	//
	////读取列表
	//values,_:=redis.Values(c.Do("lrange","redis","0","100"))
	//var s1 string
	//redis.Scan(values,&s1)
	//fmt.Println(s1)
}
func lrangeRedis(){

}
func saddRedis(){

}
func smemberSRedis(){

}