package routers

import (
	"github.com/astaxie/beego"
	"sifan/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
