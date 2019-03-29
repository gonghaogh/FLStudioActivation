package main

import (
	_"shanghaiyiqi/models"
	_ "shanghaiyiqi/routers"
	"github.com/astaxie/beego"
)

func main() {
	//绑定视图函数
	beego.AddFuncMap("prePage",ShowPrePage)
	beego.AddFuncMap("nextPage",ShowNextPage)
	beego.Run()
}
//后台定义一个函数,上一页
func ShowPrePage(pageIndex int) int{
	if pageIndex <=1{
		return 1
	}
	return pageIndex-1
}
//后台定义一个函数,下一页
func ShowNextPage(pageIndex int,totalPage int) int{
	if pageIndex>=totalPage{
		return totalPage
	}
	return pageIndex+1
}