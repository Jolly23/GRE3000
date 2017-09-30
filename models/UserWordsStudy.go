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

func BuildWordsListForUser(id int) bool {
	o := orm.NewOrm()
	// 注入前删除已有
	o.Raw("DELETE FROM user_words_study WHERE user_id = ?", id).Exec()
	var wordsList []*UserWordsStudy
	for _, eachWord := range LoadWords() {
		wordsList = append(wordsList, &UserWordsStudy{UserId: id, WordId: eachWord.Id})
	}
	o.InsertMulti(len(wordsList), wordsList)

	return true
}
