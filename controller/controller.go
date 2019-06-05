package controller

type Controller struct {
	Ct        *Context
	ChildName string
}

func (c *Controller) Init(ct *Context, cn string) {

	c.ChildName = cn
	c.Ct = ct

}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}
