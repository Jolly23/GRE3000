package controllers

import (
	"github.com/astaxie/beego"
)

type Index2Controller struct {
	beego.Controller
}

func (c *Index2Controller) Index() {

	//username := c.Ctx.Input.Param(":username")
	//ok, user := models.FindUserByUserName(username)
	//if ok {
	//	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	//	c.Data["PageTitle"] = "个人主页"
	//	c.Data["CurrentUserInfo"] = user
	//	c.Data["Topics"] = models.FindTopicByUser(&user, 7)
	//	c.Data["Replies"] = models.FindReplyByUser(&user, 7)
	//}
	c.Layout = "layout/layout.tpl"
	c.TplName = "words/test.tpl"
}
