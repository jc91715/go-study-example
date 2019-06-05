package controller

import (
	"contex"
	"fmt"
)

type Controller struct {
	Ct        *contex.Context
	ChildName string
}

func (c *Controller) Init(ct *contex.Context, cn string) {

	c.ChildName = cn
	c.Ct = ct

	fmt.Println("hello Init")

}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}
