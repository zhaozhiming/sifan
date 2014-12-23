package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int64  `orm:"auto"`
	Name string `orm:"size(128)"`
	Age  int32
}

func init() {
	orm.RegisterModel(new(User))

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	dbuser := beego.AppConfig.String("mysqluser")
	dbpass := beego.AppConfig.String("mysqlpass")
	dbhost := beego.AppConfig.String("mysqlhost")
	dbport, _ := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")

	mysql_url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)
	beego.Info(mysql_url)
	orm.RegisterDataBase("default", "mysql", mysql_url)

	orm.Debug = true
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
