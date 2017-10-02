package controllers

import (
	"GRE3000/filters"
	"GRE3000/models"
	"github.com/astaxie/beego"
	"strconv"
)

type WordsController struct {
	beego.Controller
}

func (c *WordsController) Index() {
	isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
	c.Data["IsLogin"], c.Data["UserInfo"] = isLogin, UserInfo

	var rawWordsList []*models.WordsList
	var userWordsList []*models.UserWordsStudy

	if isLogin {
		userWordsList = models.LoadWordsListForUser(&UserInfo)
	} else {
		rawWordsList = models.LoadRawWords()
	}
	if isLogin {
		c.Data["PageTitle"] = UserInfo.Username + "同学的单词表"
	} else {
		c.Data["PageTitle"] = "GRE单词表"
	}
	c.Data["RawWords"] = &rawWordsList
	c.Data["UserWords"] = &userWordsList

	c.Layout = "layout/layout.tpl"
	c.TplName = "words/vocabulary.tpl"
}

func (c *WordsController) IncrMark() {
	id := c.Ctx.Input.Param(":id")
	userWordId, _ := strconv.Atoi(id)
	if userWordId > 0 {
		isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
		if isLogin {
			userWord := models.FindUserWordByWordId(&UserInfo, userWordId)
			models.IncrWordMark(userWord, &UserInfo)
			c.Data["json"] = map[string]int{"ErrCode": 0}
			c.ServeJSON()
			return
		}
	}
	c.Data["json"] = map[string]int{"ErrCode": -1}
	c.ServeJSON()
}
