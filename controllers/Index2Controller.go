package controllers

import (
	"github.com/astaxie/beego"
)

type Index2Controller struct {
	beego.Controller
}

func (c *Index2Controller) Index() {
	c.TplName = "index.tpl"
}
