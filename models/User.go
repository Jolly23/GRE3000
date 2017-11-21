package models

import (
	"GRE3000/utils"
	"github.com/astaxie/beego/orm"
	"github.com/casbin/casbin"
	"strconv"
	"time"
)

type User struct {
	Id        int    `orm:"pk;auto;index"`
	Username  string `orm:"unique;index"`
	Password  string
	Token     string `orm:"unique;index"`
	Avatar    string
	Email     string      `orm:"null"`
	Url       string      `orm:"null"`
	Signature string      `orm:"null;size(1000)"`
	Operation []*UserLogs `orm:"reverse(many)"`
	InTime    time.Time   `orm:"auto_now_add;type(datetime)"`
	Roles     []*Role     `orm:"rel(m2m)"`
}

var Enforcer *casbin.Enforcer = nil

func getAttr(name string, attr string) string {
	if attr != "url" {
		return ""
	}

	permissions := FindPermissions()
	for _, permission := range permissions {
		if name == strconv.Itoa(permission.Id) {
			return permission.Url
		}
	}
	return ""
}

func getAttrFunc(args ...interface{}) (interface{}, error) {
	name := args[0].(string)
	attr := args[1].(string)

	return (string)(getAttr(name, attr)), nil
}

func Init() {
	Enforcer = &casbin.Enforcer{}
	Enforcer.InitWithFile("rbac_model.conf", "")

	Enforcer.AddActionAttributeFunction(getAttrFunc)

	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select user_id, role_id from user_roles").Values(&res, "user_id", "role_id")
	for _, param := range res {
		Enforcer.AddRoleForUser(param["user_id"].(string), param["role_id"].(string))
	}

	o = orm.NewOrm()
	o.Raw("select role_id, permission_id from role_permissions").Values(&res, "role_id", "permission_id")
	for _, param := range res {
		Enforcer.AddPermissionForUser(param["role_id"].(string), param["permission_id"].(string))
	}
}

func FindUserById(id int) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func Login(username string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

func PageUser(p int, size int) utils.Page {
	o := orm.NewOrm()
	var user User
	var list []User
	qs := o.QueryTable(user)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func FindPermissionByUserIdAndPermissionName(userId int, name string) bool {
	permissions := FindPermissions()
	for _, permission := range permissions {
		if name == permission.Name {
			return Enforcer.Enforce(strconv.Itoa(userId), permission.Url)
		}
	}

	return false
}

func DeleteUser(user *User) {
	Enforcer.DeleteUser(strconv.Itoa(user.Id))

	o := orm.NewOrm()
	o.Delete(user)
}

func DeleteUserRolesByUserId(UserID int) {
	Enforcer.DeleteRolesForUser(strconv.Itoa(UserID))
	o := orm.NewOrm()
	o.Raw("delete from user_roles where user_id = ?", UserID).Exec()
}

func SaveUserRole(UserID int, RoleID int) {
	Enforcer.AddRoleForUser(strconv.Itoa(UserID), strconv.Itoa(RoleID))
	o := orm.NewOrm()
	o.Raw("insert into user_roles (user_id, role_id) values (?, ?)", UserID, RoleID).Exec()
}

func FindUserRolesByUserId(UserID int) []orm.Params {
	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select id, user_id, role_id from user_roles where user_id = ?", UserID).Values(&res, "id", "user_id", "role_id")
	return res
}
