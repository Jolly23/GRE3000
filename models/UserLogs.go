package models

type UserLogs struct {
	Id      int    `orm:"pk;auto"`
	User    *User  `orm:"rel(fk)"`
	Content string `orm:"type(text)"`
}

func (u *Common) UserLogs() string {
	return "user_logs"
}
