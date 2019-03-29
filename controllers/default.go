package controllers

import (

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghaiyiqi/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"] = "china"
	c.TplName = "test.html"
}
func (c *MainController) Post(){
	c.Data["data"] = "上海一期最棒"
	c.TplName = "test.html"


}
func (c *MainController) ShowGet(){
	//获取orm对象
	o:=orm.NewOrm();
	//执行某个操作函数，增删改查
	/*
	//插入操作
	var user models.User
	user.Name = "tianqi"
	user.Password = "qing"
	id,err := o.Insert(&user)
	if err!=nil{
		beego.Error("插入错误",err)
	}
	beego.Info(id)
	*/
	/*
	//查询操作
	var user models.User
	user.Id =3
	err :=o.Read(&user,"Id")
	if err!=nil{
		beego.Error("查询错误",err)
	}
	beego.Info(user)
	*/
	//更新操作
	var user models.User
	user.Id = 3
	err := o.Read(&user)
	if err!=nil{
		beego.Error("更新的数据不存在",err)
		return
	}
	user.Name = "mingtian"
	count,err1 :=o.Update(&user)
	if err1!=nil{
		beego.Error("更新失败",err1)
	}
	beego.Info(count)
	//返回结果
	c.Data["data"] = "上海一期最棒"
	c.TplName = "test.html"
}
