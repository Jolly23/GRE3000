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
	var rawWordsList []*models.WordsList
	var userWordsList []*const_conf.RawWord

	isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
	c.Data["IsLogin"], c.Data["UserInfo"] = isLogin, UserInfo
	needRandom, _ := c.GetBool("RandomSort")

	if isLogin && !needRandom {
		userWordsList = models.LoadWordsListForUser(&UserInfo)
	} else if !needRandom {
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
	c.Data["ShowMeans"], _ = c.GetBool("ShowMeans")
	c.Data["RandomSort"] = needRandom
	c.Layout = "layout/layout.tpl"
	c.TplName = "words/vocabulary.tpl"
}

func (c *WordsController) LoadWordsJson() {
	needRandom, _ := c.GetBool("RandomSort")
	isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
	if isLogin {
		c.Data["json"] = models.LoadUserWordsJson(&UserInfo, needRandom)
	} else {
		c.Data["json"] = models.LoadRawWordsJson(needRandom)
	}
	c.ServeJSON()
}

func (c *WordsController) IncrMark() {
	ErrCode := -1
	id := c.Ctx.Input.Param(":id")
	token, flag := c.GetSecureCookie(const_conf.CookieSecure, const_conf.WebCookieName)
	userWordId, err := strconv.Atoi(id)
	if flag && !cache.Redis.IsExist(token+id) && err == nil && userWordId > 0 {
		isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
		if isLogin {
			userWord, ok := models.FindUserWordByWordId(&UserInfo, userWordId)
			if ok {
				go models.IncrWordMark(userWord, &UserInfo)
				go cache.Redis.Put(token+id, UserInfo.Username, time.Duration(const_conf.MarkWordTimeLimit)*time.Minute)
				ErrCode = 0
			}
		}
	}
	c.Data["json"] = map[string]int{"ErrCode": ErrCode}
	c.ServeJSON()
}

func (c *WordsController) DeleteWord() {
	ErrCode := -1
	id := c.Ctx.Input.Param(":id")
	userWordId, err := strconv.Atoi(id)
	if err == nil && userWordId > 0 {
		isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
		if isLogin {
			go models.DeleteWord(&UserInfo, userWordId)
			ErrCode = 0
		}
	}
	c.Data["json"] = map[string]int{"ErrCode": ErrCode}
	c.ServeJSON()
}

func (c *WordsController) Statistics() {
	isLogin, UserInfo := filters.IsLogin(c.Controller.Ctx)
	var countMarked, countAll int64
	if isLogin {
		countMarked = models.CountOfMarkedWords(&UserInfo)
		countAll = models.CountOfUserWords(&UserInfo)
	}
	c.Data["json"] = map[string]int64{"Marked": countMarked, "All": countAll}
	c.ServeJSON()
}