package routers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.html"
}


func init() {
	beego.Router("/", &MainController{})
}