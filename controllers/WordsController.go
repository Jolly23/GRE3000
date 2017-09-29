package controllers

import (
	"GRE3000/models"
	"fmt"
	"github.com/astaxie/beego"
)

type WordsController struct {
	beego.Controller
}

func (c *WordsController) Index() {
	list := models.LoadWords()

	for k, v := range list {
		fmt.Println(k, v)
	}
	c.Data["json"] = &list
	c.ServeJSON()
}
