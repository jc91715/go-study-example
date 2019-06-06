package controller

import (
	"fmt"
)

type PostController struct {
	Controller
}

func (c *PostController) Index() {
	fmt.Printf("Hello PostController Index")
}
func (c *PostController) Show() {
	fmt.Println("\nHello PostController Show:%s", c.Ct.Params["post_id"])

}
