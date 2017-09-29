package models

type UserLogs struct {
	Id   int          `orm:"pk;auto"`
	User []*UsersList `orm:"reverse(many)"`
}

func (u *Common) UserLogs() string {
	return "user_logs"
}
