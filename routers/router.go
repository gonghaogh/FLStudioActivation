package routers

import (
	"github.com/astaxie/beego/context"
	"shanghaiyiqi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/article/*",beego.BeforeExec,Filter)
	beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:DoLogin")
	beego.Router("/register", &controllers.UserController{},"get:Register;post:DoRegister")
	beego.Router("/article/logout",&controllers.UserController{},"get:Logout")
	beego.Router("/article/articleList",&controllers.ArticleController{},"get:ShowArticle")
	beego.Router("/article/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:AddArticle")
	beego.Router("/article/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")
	//编辑文章
	beego.Router("/article/updateArticleDetail",&controllers.ArticleController{},"get:UpdateArticle;post:DoUpdate")
	//删除文章
	beego.Router("/article/delArticle",&controllers.ArticleController{},"get:DelArticle")
	beego.Router("/article/articleType",&controllers.ArticleController{},"get:ShowType;post:AddType")
	beego.Router("/article/delType",&controllers.ArticleController{},"get:DelType")

    //指定自定义方法,一个请求指定一个方法
    //beego.Router("/login",&controllers.LoginController{},"get:ShowLogin;post:PostFunc")
    ////给多个请求指定一个方法
	//beego.Router("/index",&controllers.IndexController{},"get,post:HandleFunc")
    ////给所有请求指定一个方法
    //beego.Router("/index",&controllers.IndexController{},"*:HandleFunc;")
	////当两种指定方法冲突的时候
	//beego.Router("/index",&controllers.IndexController{},"*:HandleFunc;post:PostFunc")
	//beego.Router("/", &controllers.MainController{},"get:ShowGet;post:Post")
}
var Filter = func(ctx *context.Context) {
	userName :=ctx.Input.Session("userName")
	if userName ==nil{
		ctx.Redirect(302,"/login")
		return
	}
}
