package controllers

import (
	"beego-bbs/models/post"
	"github.com/astaxie/beego"
)

type PostController struct {
	beego.Controller
	repository post.PostRepository
}

func (this *PostController) Index() {
	posts, _ := this.repository.FindAll()
	this.Data["posts"] = posts
	this.Layout = "layouts/application.html"
	this.TplNames = "post/index.html"
}

func (this *PostController) Create() {
	post := post.Post{
		Content: this.GetString("content"),
	}

	err := this.repository.Save(&post)
	flash := beego.NewFlash()
	if err != nil {
		flash.Error("The post could not be saved. Please, try again.")
	} else {
		flash.Notice("The post has been saved.")
	}
	flash.Store(&this.Controller)
	
	this.Redirect("/", 302)
}
