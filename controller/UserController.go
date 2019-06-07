package controller

import (
	"fmt"
	"model"
	"net/http"
	"strconv"
)

type UserController struct {
	Controller
}

func (c *UserController) Index() {
	fmt.Printf("\nHello UserController Index")
}

func (c *UserController) Show() {
	userModel := model.NewUserModel()

	// defer c.userModel.Db.Close()

	id, ok := c.Ct.Params["user_id"] /*如果确定是真实的,则存在,否则不存在 */

	if ok {
		intId, err := strconv.Atoi(id)
		CheckErr(err)
		u := userModel.Find(intId)
		fmt.Println(u)
		m := u.Find(5)
		fmt.Println(m)
	} else {
		fmt.Println("不存在")
		http.NotFound(c.Ct.ResponseWriter, c.Ct.Request)
	}
	fmt.Printf("\nHello UserController Show:%s", id)
}
func (c *UserController) Prepare() {

	fmt.Println("\nhello Prepare")
}
