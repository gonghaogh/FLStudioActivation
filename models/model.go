package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
//表的设计
//定义一个结构体
type User struct {
	Id int
	Name string
	Password string

	Article []*Article `orm:"reverse(many)"`
}
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string `orm:"size(100)"`

	ArticleType *ArticleType `orm:"rel(fk)"`
	Users []*User `orm:"rel(m2m)"`
}
//类型表
type ArticleType struct{
	Id int
	TypeName string `orm:"size(20)"`

	Article []*Article `orm:"reverse(many)"`
}


func init(){

	//ORM操作数据库

	//获取连接对象
	orm.RegisterDataBase("default","mysql","root:1@tcp(127.0.0.1:3306)/test?charset=utf8")
	//创建表
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	//生成表
	//第一个参数是数据库别名，第二个参数是是否强制更新，第三个参数过程是否可见
	orm.RunSyncdb("default",false,true)

	/*
	//操作数据库代码
	//第一个参数是数据库驱动
	//链接数据库字符串
	//root:123456@tcp(127.0.0.1:3306)
	conn,err := sql.Open("mysql","root:1@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err !=nil{
		beego.Info("链接错误",err)
		beego.Error("链接错误",err)
		return
	}
	//关闭数据库
	defer conn.Close()
	//创建表
	//_,err1 := conn.Exec("create table user(name VARCHAR(40),password VARCHAR(40))")
	//if err1 !=nil{
	//	beego.Info("创建错误",err1)
	//	beego.Error("创建错误",err1)
	//	return
	//}
	//插入数据
	//conn.Exec("insert into user(name,password) values(?,?)","今天","晴")

	//查询
	res,err:=conn.Query("select name from user")
	var name string
	for res.Next(){
		res.Scan(&name)
		beego.Info(name)
	}
	*/


}