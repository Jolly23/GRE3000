package controllers

import (
	"GRE3000/const_conf"
	"github.com/astaxie/beego"
)

type CookieController struct {
	beego.Controller
}

func (c *CookieController) Test() {
	c.SetSecureCookie(const_conf.CookieSecure, "Naaaame", "hahaha", const_conf.OpenidCookieExpire, "/", const_conf.DomainName)
	c.Data["vvv"] = "设置cookie成功"
	c.TplName = "test.tpl"
}

func (c *CookieController) Test2() {
	token, flag := c.GetSecureCookie(const_conf.CookieSecure, "Naaaame")

	if flag {
		c.Data["vvv"] = token
	} else {
		c.Data["vvv"] = "没找到"
	}
	c.TplName = "test.tpl"
}
