package controllers

import (
	"GRE3000/models"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Index() {
	models.BuildWordsListForUser("jolly")
	c.Ctx.WriteString("OK")
}
