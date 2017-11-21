package controllers

import (
	"github.com/Jolly23/GRE3000/const_conf"
	"github.com/Jolly23/GRE3000/filters"
	"github.com/Jolly23/GRE3000/models"
	"github.com/astaxie/beego"
	"net/http"
	"regexp"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Detail() {
	username := c.Ctx.Input.Param(":username")
	ok, user := models.FindUserByUserName(username)
	if ok {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
		c.Data["PageTitle"] = "个人主页"
		c.Data["CurrentUserInfo"] = user
		c.Data["Topics"] = models.FindTopicByUser(&user, 7)
		c.Data["Replies"] = models.FindReplyByUser(&user, 7)
	}
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

func (c *UserController) ToSetting() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["PageTitle"] = "用户设置"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}

func (c *UserController) Setting() {
	flash := beego.NewFlash()
	email, url, signature := c.Input().Get("email"), c.Input().Get("url"), c.Input().Get("signature")
	if len(email) > 0 {
		ok, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|_|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", email)
		if !ok {
			flash.Error("请输入正确的邮箱地址")
			flash.Store(&c.Controller)
			c.Redirect("/user/setting", 302)
			return
		}
	}
	if len(signature) > 1000 {
		flash.Error("个人签名长度不能超过1000字符")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	_, user := filters.IsLogin(c.Ctx)
	user.Email = email
	user.Url = url
	user.Signature = signature
	go models.UpdateUser(&user)
	flash.Success("更新资料成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *UserController) UpdatePwd() {
	flash := beego.NewFlash()
	oldPwd, newPwd := c.Input().Get("oldpwd"), c.Input().Get("newpwd")
	_, user := filters.IsLogin(c.Ctx)
	if user.Password != oldPwd {
		flash.Error("旧密码不正确")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	if len(newPwd) == 0 {
		flash.Error("新密码都不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	user.Password = newPwd
	go models.UpdateUser(&user)
	flash.Success("密码修改成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *UserController) UpdateAvatar() {
	flash := beego.NewFlash()
	f, h, err := c.GetFile("avatar")
	if err == http.ErrMissingFile {
		flash.Error("请选择文件")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	}
	defer f.Close()
	if err != nil {
		flash.Error("上传失败")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	} else {
		c.SaveToFile("avatar", "static/upload/avatar/"+h.Filename)
		_, user := filters.IsLogin(c.Ctx)
		user.Avatar = "/static/upload/avatar/" + h.Filename
		go models.UpdateUser(&user)
		flash.Success("上传成功")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	}
}

func (c *UserController) List() {
	c.Data["PageTitle"] = "用户列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	c.Data["Page"] = models.PageUser(p, const_conf.PageSize)
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/list.tpl"
}

func (c *UserController) Edit() {
	c.Data["PageTitle"] = "配置角色"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		ok, user := models.FindUserById(id)
		if ok {
			c.Data["User"] = user
			c.Data["Roles"] = models.FindRoles()
			c.Data["UserRoles"] = models.FindUserRolesByUserId(id)
			c.Layout = "layout/layout.tpl"
			c.TplName = "user/edit.tpl"
		} else {
			c.Ctx.WriteString("用户不存在")
		}
	} else {
		c.Ctx.WriteString("用户不存在")
	}
}

func (c *UserController) Update() {
	c.Data["PageTitle"] = "配置角色"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	roleIds := c.GetStrings("roleIds")
	if id > 0 {
		go models.DeleteUserRolesByUserId(id)
		for _, v := range roleIds {
			roleId, _ := strconv.Atoi(v)
			models.SaveUserRole(id, roleId)
		}
		c.Redirect("/user/list", 302)
	} else {
		c.Ctx.WriteString("用户不存在")
	}
}

func (c *UserController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		ok, user := models.FindUserById(id)
		if ok {
			go models.DeleteTopicByUser(&user)
			go models.DeleteReplyByUser(&user)
			go models.DeleteWordsListForUser(&user)
			go models.DeleteUser(&user)
		}
		c.Redirect("/user/list", 302)
	} else {
		c.Ctx.WriteString("用户不存在")
	}
}
