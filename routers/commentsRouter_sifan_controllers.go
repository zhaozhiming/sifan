package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["sifan/controllers:MainController"] = append(beego.GlobalControllerRouter["sifan/controllers:MainController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sifan/controllers:MenuController"] = append(beego.GlobalControllerRouter["sifan/controllers:MenuController"],
		beego.ControllerComments{
			"AddUser",
			`/user/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sifan/controllers:MenuController"] = append(beego.GlobalControllerRouter["sifan/controllers:MenuController"],
		beego.ControllerComments{
			"FindUser",
			`/user/:id`,
			[]string{"get"},
			nil})

}
