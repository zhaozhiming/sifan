package controllers

import (
	"github.com/astaxie/beego"
	"sifan/models"
)

// oprations for Menu
type MenuController struct {
	beego.Controller
}

func (c *MenuController) URLMapping() {
	c.Mapping("AddUser", c.AddUser)
	c.Mapping("FindUser", c.FindUser)
}

// @router /user/add [post]
func (c *MenuController) AddUser() {
	beego.Info("add user...")

	user := new(models.User)
	user.Name = c.GetString("name")
	age, _ := c.GetInt("age")
	user.Age = int32(age)
	models.AddUser(user)

	c.Data["json"] = &user
	c.ServeJson()
}

// @router /user/:id [get]
func (c *MenuController) FindUser() {
	beego.Info("get user...")
	id, _ := c.GetInt(":id")
	user, _ := models.GetUserById(int64(id))
	c.Data["json"] = &user
	c.ServeJson()
}
