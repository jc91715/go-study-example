package controller

import (
	"fmt"
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

	// fmt.Printf(u.Ct.Value())
	fmt.Printf("Hello UserController Show")
}
