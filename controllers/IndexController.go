package controllers

import (
	"GRE3000/const_conf"
	"GRE3000/filters"
	"GRE3000/models"
	"github.com/astaxie/beego"
	"github.com/sluu99/uuid"
	"strconv"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Prepare() {
	c.EnableXSRF = false
}

//首页
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	s, _ := strconv.Atoi(c.Ctx.Input.Query("s"))
	c.Data["S"] = s
	section := models.Section{Id: s}
	c.Data["Page"] = models.PageTopic(p, const_conf.PageSize, &section)
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
}

//登录页
func (c *IndexController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "登录"
		c.Layout = "layout/layout.tpl"
		c.TplName = "login.tpl"
	}
}

//验证登录
func (c *IndexController) Login() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")
	if flag, user := models.Login(username, password); flag {
		c.SetSecureCookie(const_conf.CookieSecure, const_conf.WebCookieName, user.Token, const_conf.CookieExpire, "/", const_conf.DomainName, false, true)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

//注册页
func (c *IndexController) RegisterPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "注册"
		c.Layout = "layout/layout.tpl"
		c.TplName = "register.tpl"
	}
}

//验证注册
func (c *IndexController) Register() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")
	if len(username) == 0 || len(password) == 0 {
		flash.Error("用户名或密码不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if flag, _ := models.FindUserByUserName(username); flag {
		flash.Error("用户名已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		var token = uuid.Rand().Hex()
		user := models.User{Username: username, Password: password, Avatar: "/static/imgs/avatar.png", Token: token}
		new_user_id := models.SaveUser(&user)
		models.SaveUserRole(int(new_user_id), 5)
		//models.BuildWordsListForUser(int(new_user_id))

		// others are ordered as cookie's max age time, path,domain, secure and http only.
		c.SetSecureCookie(const_conf.CookieSecure, const_conf.WebCookieName, token, const_conf.CookieExpire, "/", const_conf.DomainName, false, true)
		c.Redirect("/", 302)
	}
}

//登出
func (c *IndexController) Logout() {
	c.SetSecureCookie(const_conf.CookieSecure, const_conf.WebCookieName, "", -1, "/", const_conf.DomainName, false, true)
	c.Redirect("/", 302)
}

//关于
func (c *IndexController) About() {
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "关于"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}
