package models

import (
	"time"
)

type UsersList struct {
	Id        int    `orm:"pk;auto;index"`
	Username  string `orm:"size(256);unique;index"`
	Password  string `orm:"size(256);"`
	Token     string `orm:"unique;index"`
	Avatar    string
	Email     string    `orm:"null"`
	Url       string    `orm:"null"`
	Signature string    `orm:"null;size(1000)"`
	Operation *UserLogs `orm:"rel(fk)"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
}

func (u *Common) UsersList() string {
	return "users_list"
}
