package controllers

import (
	"GRE3000/base/cache"
	"GRE3000/const_conf"
	"GRE3000/filters"
	"GRE3000/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
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
	c.Data["IsWordsPage"] = true
	c.Data["RawWords"] = &rawWordsList
	c.Data["UserWords"] = &userWordsList

	c.Layout = "layout/layout.tpl"
	c.TplName = "words/vocabulary.tpl"
}

func (c *WordsController) IncrMark() {
	id := c.Ctx.Input.Param(":id")
	token, flag := c.GetSecureCookie(const_conf.CookieSecure, const_conf.WebCookieName)
	cacheKeyName := token + id
	if flag {
		val := cache.Redis.Get(cacheKeyName)
		if val != nil {
			c.Data["json"] = map[string]int{"ErrCode": -1}
			c.ServeJSON()
			return
		}
	}

	userWordId, _ := strconv.Atoi(id)
	if userWordId > 0 {
		isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
		if isLogin {
			userWord := models.FindUserWordByWordId(&UserInfo, userWordId)
			models.IncrWordMark(userWord, &UserInfo)
			cache.Redis.Put(cacheKeyName, UserInfo.Username, time.Duration(const_conf.MarkWordTimeLimit)*time.Minute)
			c.Data["json"] = map[string]int{"ErrCode": 0}
			c.ServeJSON()
			return
		}
	}
	c.Data["json"] = map[string]int{"ErrCode": -1}
	c.ServeJSON()
}

func (c *WordsController) DeleteWord() {
	id := c.Ctx.Input.Param(":id")
	userWordId, _ := strconv.Atoi(id)
	if userWordId > 0 {
		isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
		if isLogin {
			models.DeleteWord(&UserInfo, userWordId)
			c.Data["json"] = map[string]int{"ErrCode": 0}
			c.ServeJSON()
		}
	}
	c.Data["json"] = map[string]int{"ErrCode": -1}
	c.ServeJSON()
}
