package models

import (
	"GRE3000/const_conf"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strconv"
)

type UserLogs struct {
	Id      int    `orm:"pk;auto"`
	User    *User  `orm:"rel(fk);index"`
	Content string `orm:"type(text);index"`
}

func (u *Common) UserLogs() string {
	return "user_logs"
}

func NewLog(newLog *UserLogs) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(newLog)
	return id
}

func FindUserLastMarkWord(user *User) (string, bool) {
	o := orm.NewOrm()
	var log UserLogs
	err := o.QueryTable(log).Filter("user_id", user.Id).Filter("Content__startswith", const_conf.LogMarkWord).OrderBy("-id").Limit(1).One(&log)
	if err != orm.ErrNoRows {
		reg := regexp.MustCompile(`[\d]+`)
		slice := reg.FindAllString(log.Content, -1)
		num, err := strconv.Atoi(slice[0])
		if err == nil {
			word := FindWordById(num)
			return word.Word, true
		}
	}
	return "", false
}
