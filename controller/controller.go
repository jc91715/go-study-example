package controller

import (
	"context"
)

type Controller struct {
	Ct        *context.Context
	ChildName string
}

func (c *Controller) Init(ct *context.Context, cn string) {

	c.ChildName = cn
	c.Ct = ct

}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}
func (c *Controller) RouteMaps() {

}
