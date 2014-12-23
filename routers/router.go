package routers

import (
	"github.com/astaxie/beego"
	"sifan/controllers"
)

func init() {
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.MenuController{})
}
