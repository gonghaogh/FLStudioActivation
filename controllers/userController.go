package controllers

import (
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghaiyiqi/models"
)

type UserController struct {
	beego.Controller
}
func(this *UserController) ShowLogin(){
	userName := this.Ctx.GetCookie("userName")
	if userName != ""{
		temp,err := base64.StdEncoding.DecodeString(userName)
		if err!=nil{
			beego.Error("base64.StdEncoding.DecodeString err",err)
			this.Data["userName"]=""
			this.Data["checked"] =""
		}else{
			this.Data["userName"]=string(temp)
			this.Data["checked"] ="checked"
		}
	}else{
		this.Data["userName"]=""
		this.Data["checked"] =""
	}
	this.TplName="login.html"
}
func (this *UserController) Logout(){
	this.DelSession("userName")
	this.Redirect("/login",302)
}
func(this *UserController) DoLogin(){
	//获取数据
	userName :=this.GetString("userName")
	password := this.GetString("password")
	remember :=this.GetString("remember")
	if remember == "on"{
		temp := base64.StdEncoding.EncodeToString([]byte(userName))
		this.Ctx.SetCookie("userName",temp,1000)
	}else{
		this.Ctx.SetCookie("userName",userName,-1)
	}
	//beego.Info(userName,password)
	//校验数据
	if userName=="" || password ==""{
		this.Data["errmsg"] = "登录数据不完整，请重新登录"
		this.TplName = "login.html"
		return
	}
	//验证用户是否存在
	o:=orm.NewOrm()
	var user models.User
	//user.Password = password
	user.Name = userName
	err := o.Read(&user,"name")
	beego.Info(user)
	if err!=nil {
		this.Data["errmsg"] = "用户名或密码错误，请重新登录"
		this.TplName = "login.html"
		return
	}
	if user.Password == password {
		this.SetSession("userName",userName)
		this.Redirect("/articleList",302)
	}else{
		this.Data["errmsg"] = "用户名或密码错误，请重新登录"
		this.TplName = "login.html"
		return
	}
}


func(u *UserController) Register(){
	u.TplName = "register.html"

}
func(u *UserController) DoRegister(){
	//获取数据
	userName :=u.GetString("userName")
	password := u.GetString("password")
	//beego.Info(userName,password)
	//校验数据
	if userName=="" || password ==""{
		u.Data["errmsg"] = "注册数据不完整，请重新注册"
		u.TplName = "register.html"
		return
	}
	//操作数据
	o :=orm.NewOrm()
	var user models.User
	user.Name=userName
	user.Password = password
	_,err :=o.Insert(&user)
	if err!=nil{
		beego.Error("插入数据错误",err)
	}else{
		//跳转到登录页面
		u.Redirect("/login",302)
	}
}