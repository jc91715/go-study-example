package controller

import (
	"fmt"
	"net/http"
)

type UserController struct {
	Controller
}

func (c *Controller) Get() {
	fmt.Printf("Hello UserController Get")
}
func (c *Controller) Index() {
	fmt.Printf("Hello UserController Index")
}

func (c *Controller) Show() {

	id, ok := c.Ct.Params["id"] /*如果确定是真实的,则存在,否则不存在 */

	if ok {
		fmt.Println("ok", id)
	} else {
		fmt.Println("不存在")
		http.NotFound(c.Ct.ResponseWriter, c.Ct.Request)
	}
	fmt.Printf("Hello UserController Show")
	fmt.Printf("Hello UserController Show")
}
