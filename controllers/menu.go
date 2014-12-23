package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"sifan/models"
)

// oprations for Menu
type MenuController struct {
	beego.Controller
}

func (c *MenuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Menu
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 200 {int} models.Menu.Id
// @Failure 403 body is empty
// @router / [post]
func (c *MenuController) Post() {
	beego.Info("post")

	user := new(models.User)
	user.Name = "zzm"
	user.Age = 18
	models.AddUser(user)

	mystruct := map[string]string{"result": "success"}
	c.Data["json"] = &mystruct
	c.ServeJson()
}

// @Title Get
// @Description get Menu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Menu
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MenuController) GetOne() {
	fmt.Println("hello menu")
}

func (c *MenuController) Get() {
	beego.Info("hello")
	mystruct := map[string]string{"result": "success"}
	c.Data["json"] = &mystruct
	c.ServeJson()
}

// @Title Get All
// @Description get Menu
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Menu
// @Failure 403
// @router / [get]
func (c *MenuController) GetAll() {
	fmt.Println("hello menu")
}

// @Title Update
// @Description update the Menu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 200 {object} models.Menu
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MenuController) Put() {

}

// @Title Delete
// @Description delete the Menu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MenuController) Delete() {

}
