package controller

import (
	"fmt"
	"model"

	// "html/template"
	"net/http"
	// "os"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type UserController struct {
	Controller
}

func (c *UserController) Index() {
	fmt.Printf("\nHello UserController Index")
}

func (c *UserController) Show() {
	// userModel := model.NewUserModel()

	// // defer c.userModel.Db.Close()

	id, ok := c.Ct.Params["user_id"] /*如果确定是真实的,则存在,否则不存在 */

	// if ok {
	// 	intId, err := strconv.Atoi(id)
	// 	CheckErr(err)
	// 	u := userModel.Find(intId)
	// 	// fmt.Println(u)
	// 	m := u.Find(5)
	// 	fmt.Println(m)
	// } else {
	// 	fmt.Println("不存在")
	// 	http.NotFound(c.Ct.ResponseWriter, c.Ct.Request)
	// }
	// fmt.Printf("\nHello UserController Show:%s", id)

	if !ok {
		fmt.Println("不存在")
		http.NotFound(c.Ct.ResponseWriter, c.Ct.Request)
	} else {
		o := orm.NewOrm()
		intId, err := strconv.Atoi(id)
		CheckErr(err)
		user := model.Userinfo{Uid: intId}
		err = o.Read(&user)
		if err == orm.ErrNoRows {
			fmt.Println("查询不到")
		} else if err == orm.ErrMissPK {
			fmt.Println("找不到主键")
		} else {
			fmt.Println(user.Uid, user.Username, user.Department)
		}

		// s1.Execute(c.Ct.ResponseWriter, nil)
	}

}
func (c *UserController) Prepare() {

	fmt.Println("\nhello Prepare")
}
