package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"shanghaiyiqi/models"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowArticle(){
	session := this.GetSession("userName")
	if session ==nil{
		this.Redirect("/login",302)
		return
	}

	//获取当前页
	pageIndex,err2 := this.GetInt("pageIndex")
	if err2!=nil||pageIndex<1{
		pageIndex=1
	}
	//每页显示条数
	var pageSize int = 6
	start := (pageIndex-1)*pageSize
	this.TplName="index.html"
	//查询所有的article数据
	o :=orm.NewOrm()
	qs :=o.QueryTable("article")
	var articles []models.Article
	//qs.All(&articles)
	//查询总记录数
	count,err :=qs.Count()
	if err!=nil{
		beego.Error("查询articles记录数出错")
		return;
	}
	totalPage := math.Ceil(float64(count)/float64(pageSize))

	//分页查询数据
	qs.Limit(pageSize,start).All(&articles)
	this.Data["articles"] = articles
	this.Data["totalPage"] = int(totalPage)
	this.Data["count"] = count
	this.Data["pageIndex"]=pageIndex
}
func (this *ArticleController) ShowAddArticle(){
	o := orm.NewOrm()
	var types []models.ArticleType
	qs :=o.QueryTable("article_type")
	qs.All(&types)
	this.Data["types"]=types


	this.TplName="add.html"
}
//上传文件
func uploadFile(this *beego.Controller,filepath string) string{
	//文件上传
	file,head,err := this.GetFile(filepath)
	if head.Filename ==""{
		return "NoImg"
	}
	if err!=nil{
		beego.Error("上传错误",err)
		this.Data["errmsg"] = "上传错误，请重试"
		return ""
	}
	defer file.Close()
	//判断文件大小，校验
	if head.Size > 5000000{
		this.Data["errmsg"] = "上传图片过大，请重试"
		return ""
	}
	//判断文件格式
	ext :=path.Ext(head.Filename)
	if ext!=".jpg"&&ext!=".png"&&ext!=".jpeg"{
		this.Data["errmsg"] = "不支持的图片格式，请重试"
		return ""
	}
	//以当前时间生成图片名称
	fileName := time.Now().Format("2006-01-02-15:04:05")+ext
	err1 := this.SaveToFile("uploadname","./static/img/"+fileName)
	if err1!=nil{
		return ""
	}
	return "/static/img/"+fileName
}

func (this *ArticleController) AddArticle(){
	//获取数据
	artiName := this.GetString("articleName")
	aContent := this.GetString("content")
	if artiName==""||aContent==""{
		this.Data["errmsg"] = "数据不完整，请完成添加"
		this.TplName="add.html"
		return
	}
	//文件上传
	fileName :=uploadFile(&this.Controller,"uploadname")
	//保存数据
	var article models.Article
	article.ArtiName = artiName
	article.Acontent = aContent
	if fileName != "NoImg"{
		article.Aimg = fileName
	}
	o := orm.NewOrm()
	_,err2 := o.Insert(&article)
	if err2!=nil{
		beego.Error("数据保存错误",err2)
		this.Data["errmsg"] = "系统繁忙，请重试"
		this.TplName="add.html"
		return
	}
	this.Redirect("/articleList",302)

}
func (this *ArticleController) ShowArticleDetail(){
	articleId,err := this.GetInt("articleId")
	if err != nil{
		beego.Error("this.GetInt() error",err)
		return
	}
	if articleId ==0{
		beego.Error("未获取到articleId")
		return
	}
	//通过id查询数据
	var article models.Article
	article.Id = articleId
	//查询
	o := orm.NewOrm()
	err1 := o.Read(&article)
	if err1 != nil{
		beego.Error("查询数据错误",err1)
		return
	}
	//阅读次数加1
	article.Acount +=1
	o.Update(&article)
	this.Data["article"]=article
	this.TplName = "content.html"

}
//更新数据显示
func (this *ArticleController) UpdateArticle(){
	articleId,err := this.GetInt("articleId")
	if err != nil{
		beego.Error("this.GetInt() error",err)
		return
	}
	if articleId ==0{
		beego.Error("未获取到articleId")
		return
	}
	//通过id查询数据
	var article models.Article
	article.Id = articleId
	//查询
	o := orm.NewOrm()
	err1 := o.Read(&article)
	if err1 != nil{
		beego.Error("查询数据错误",err1)
		return
	}
	this.Data["article"]=article
	this.TplName = "update.html"

}
func (this *ArticleController)  DoUpdate(){
	id,err := this.GetInt("articleId")
	if err != nil{
		beego.Error("this.GetInt() error",err)
		return
	}
	//上传文件
	filepath :=uploadFile(&this.Controller,"uploadname")

	name := this.GetString("articleName")
	content := this.GetString("content")
	var article models.Article
	if filepath != "NoImg"{
		article.Aimg = filepath
	}

	article.Id =id
	if name!=""{
		article.ArtiName=name
	}
	if content!=""{
		article.Acontent=content
	}
	article.Aimg=filepath
	o := orm.NewOrm()
	_,err2 :=o.Update(&article)
	if err2 !=nil{
		beego.Error("更新数据失败",err2)
	}
	this.Redirect("/articleList",302)

}
func (this *ArticleController) DelArticle(){
	id,err := this.GetInt("articleId")
	if err!=nil{
		beego.Error("this.GetInt err",err)
		return
	}
	var article models.Article
	article.Id = id
	o := orm.NewOrm()
	_,err2 :=o.Delete(&article)
	if err2!=nil{
		beego.Error("o.Delete err",err2)
		return
	}
	this.Redirect("/articleList",302)

}
func (this *ArticleController) ShowType(){
	var types []models.ArticleType
	o:=orm.NewOrm()
	qs :=o.QueryTable("article_type")
	qs.All(&types)
	this.Data["types"]=types
	this.TplName="addType.html"

}
func (this *ArticleController) AddType(){
	typeName := this.GetString("typeName")
	if typeName ==""{
		beego.Error("typeName为空")
	}
	o := orm.NewOrm()
	var arType models.ArticleType
	arType.TypeName = typeName
	_,err :=o.Insert(&arType)
	if err!=nil{
		beego.Error("插入articleType数据错误",err)
	}
	this.Redirect("/articleType",302)
}
func (this *ArticleController) DelType(){
	id,err := this.GetInt("id")
	if err !=nil{
		beego.Error("this.GetInt err",err)
		return
	}
	var artType models.ArticleType
	artType.Id=id
	o := orm.NewOrm()
	_,err2 :=o.Delete(&artType)
	if err2!=nil{
		beego.Error("o.Delete err",err2)
		return
	}
	this.Redirect("/articleType",302)


}
