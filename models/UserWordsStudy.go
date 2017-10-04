package models

import (
	"GRE3000/const_conf"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserWordsStudy struct {
	Id         int        `orm:"pk;auto"`
	UserId     int        `orm:"index"`
	Word       *WordsList `orm:"rel(fk)"`
	CountMarks int        `orm:"default(0);index"`
	LastMark   time.Time  `orm:"auto_now;type(datetime)"`
}

func (u *Common) UserWordsStudy() string {
	return "words_study"
}

func BuildWordsListForUser(id int) bool {
	o := orm.NewOrm()
	o.Raw("DELETE FROM user_words_study WHERE user_id = ?", id).Exec()
	var wordsList []*UserWordsStudy
	for _, eachWord := range LoadRawWords() {
		wordsList = append(wordsList, &UserWordsStudy{UserId: id, Word: eachWord})
	}
	o.InsertMulti(len(wordsList), wordsList)

	return true
}

func DeleteWordsListForUser(user *User) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM user_words_study WHERE user_id = ?", user.Id).Exec()
}

func LoadWordsListForUser(user *User) []*UserWordsStudy {
	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("SELECT user_id FROM user_words_study WHERE user_id = ? LIMIT 1", user.Id).Values(&res, "user_id")
	if len(res) == 0 {
		BuildWordsListForUser(user.Id)
	}
	var study UserWordsStudy
	var list []*UserWordsStudy
	qs := o.QueryTable(study)
	qs = qs.Filter("UserId", user.Id).RelatedSel()
	qs.OrderBy("-CountMarks").Limit(-1).All(&list)
	return list
}

func FindUserWordByWordId(user *User, wordId int) *UserWordsStudy {
	o := orm.NewOrm()
	var userWord UserWordsStudy
	o.QueryTable(userWord).Filter("UserId", user.Id).Filter("word_id", wordId).One(&userWord)
	return &userWord
}

func IncrWordMark(UserWord *UserWordsStudy, user *User) {
	o := orm.NewOrm()
	UserWord.CountMarks = UserWord.CountMarks + 1
	o.Update(UserWord, "CountMarks", "LastMark")
	o.Insert(&UserLogs{User: user, Content: fmt.Sprintf(const_conf.LogMarkWordFmt, UserWord.Word.Id)})
}

func DeleteWord(user *User, wordId int) {
	o := orm.NewOrm()
	var userWord UserWordsStudy
	o.QueryTable(userWord).Filter("UserId", user.Id).Filter("word_id", wordId).Delete()
}
