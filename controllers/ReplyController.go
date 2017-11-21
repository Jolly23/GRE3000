package controllers

import (
	"GRE3000/filters"
	"GRE3000/models"
	"GRE3000/utils"
	"github.com/astaxie/beego"
	"strconv"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Save() {
	content := c.Input().Get("content")
	if len(content) == 0 {
		c.Ctx.WriteString("回复内容不能为空")
	} else {
		tid, _ := strconv.Atoi(c.Input().Get("tid"))
		if tid == 0 {
			c.Ctx.WriteString("回复的话题不存在")
		} else {
			_, user := filters.IsLogin(c.Ctx)
			topic := models.FindTopicById(tid)
			reply := models.Reply{Content: content, Topic: &topic, User: &user, Up: 0}
			go models.SaveReply(&reply)
			go models.IncrReplyCount(&topic)
			c.Redirect("/topic/"+strconv.Itoa(tid), 302)
		}
	}
}

func (c *ReplyController) Up() {
	rid, _ := strconv.Atoi(c.Ctx.Input.Query("rid"))
	result := utils.Result{Code: 200, Description: "成功"}
	if rid > 0 {
		_, user := filters.IsLogin(c.Ctx)
		reply := models.FindReplyById(rid)
		replyUpLog := models.FindReplyUpLogByUserAndReply(&user, &reply)
		if replyUpLog.Id > 0 {
			result.Code = 201
			result.Description = "你已经赞过这条回复了"
		} else {
			replyUpLog.User = &user
			replyUpLog.Reply = &reply
			go models.SaveReplyUpLog(&replyUpLog)
			go models.UpReply(&reply)
		}
	} else {
		result.Code = 201
		result.Description = "失败1"
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *ReplyController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		reply := models.FindReplyById(id)
		tid := reply.Topic.Id
		go models.ReduceReplyCount(reply.Topic)
		go models.DeleteReply(&reply)
		c.Redirect("/topic/"+strconv.Itoa(tid), 302)
	} else {
		c.Ctx.WriteString("回复不存在")
	}
}
