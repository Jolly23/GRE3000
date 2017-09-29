package models

import (
	"github.com/astaxie/beego/orm"
)

type UserWordsStudy struct {
	Id         int `orm:"pk;auto"`
	UserId     int `orm:"index"`
	WordId     int `orm:"index"`
	CountMarks int `orm:"default(0);index"`
}

func (u *Common) UserWordsStudy() string {
	return "words_study"
}

func BuildWordsListForUser(username string) bool {
	o := orm.NewOrm()

	var tableOfUser UsersList
	var currentUser UsersList
	err := o.QueryTable(tableOfUser).Filter("Username", username).One(&currentUser)

	if err == orm.ErrNoRows {
		return false
	}

	var wordsList []*UserWordsStudy
	for _, eachWord := range LoadWords() {
		wordsList = append(wordsList, &UserWordsStudy{UserId: currentUser.Id, WordId: eachWord.Id})
	}

	o.InsertMulti(len(wordsList), wordsList)

	return true
}
