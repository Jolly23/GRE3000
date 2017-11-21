package filters

import (
	"GRE3000/const_conf"
	"GRE3000/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strconv"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
	token, flag := ctx.GetSecureCookie(const_conf.CookieSecure, const_conf.WebCookieName)
	var user models.User
	if flag {
		flag, user = models.FindUserByToken(token)
	}
	return flag, user
}

var HasPermission = func(ctx *context.Context) {
	ok, user := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/login")
	} else {
		url := ctx.Request.RequestURI
		beego.Debug("url: ", url)
		flag := models.Enforcer.Enforce(strconv.Itoa(user.Id), url)
		if !flag {
			ctx.WriteString("你没有权限访问这个页面")
		}
	}
}

var FilterUser = func(ctx *context.Context) {
	ok, _ := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}
