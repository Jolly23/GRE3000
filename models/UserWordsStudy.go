package models

import (
	"GRE3000/const_conf"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/xeonx/timeago"
	"time"
)

type UserWordsStudy struct {
	Id         int        `orm:"pk;auto;index"`
	UserId     int        `orm:"index"`
	Word       *WordsList `orm:"rel(fk)"`
	CountMarks int        `orm:"default(0);index"`
	LastMark   time.Time  `orm:"auto_now;type(datetime);null"`
}

func (u *Common) UserWordsStudy() string {
	return "words_study"
}

func BuildWordsListForUser(id int) bool {
	var wordsList []*UserWordsStudy
	var tableOfWords WordsList
	var allWords []*WordsList

	o := orm.NewOrm()
	o.Raw("DELETE FROM user_words_study WHERE user_id = ?", id).Exec()
	o.QueryTable(tableOfWords).OrderBy("id").Limit(-1).All(&allWords)
	for _, eachWord := range allWords {
		wordsList = append(wordsList, &UserWordsStudy{UserId: id, Word: eachWord})
	}
	o.InsertMulti(len(wordsList), wordsList)
	return true
}

func DeleteWordsListForUser(user *User) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM user_words_study WHERE user_id = ?", user.Id).Exec()
}

func LoadWordsListForUser(user *User) []*const_conf.RawWord {
	var userWord UserWordsStudy
	var userWords []*const_conf.RawWord

	o := orm.NewOrm()
	count, _ := o.QueryTable(userWord).Filter("UserId", user.Id).Count()
	if count == 0 {
		BuildWordsListForUser(user.Id)
	}
	o.Raw("SELECT T0.id, T1.word, T1.means, T0.count_marks, T0.last_mark FROM user_words_study T0 INNER JOIN words_list T1 ON T1.id = T0.word_id WHERE T0.user_id = ? ORDER BY T0.count_marks DESC, T1.id LIMIT ?", user.Id, const_conf.SyncLoadOffset).QueryRows(&userWords)
	return userWords
}

func LoadUserWordsJson(user *User, random bool) []*const_conf.UserWordsJson {
	var userWord UserWordsStudy
	var returnUserWords []*const_conf.UserWordsJson
	var tempWord *const_conf.UserWordsJson
	var userWords []*const_conf.RawWord

	o := orm.NewOrm()
	count, _ := o.QueryTable(userWord).Filter("UserId", user.Id).Count()
	if count == 0 {
		BuildWordsListForUser(user.Id)
	}
	if random {
		o.Raw("SELECT T0.id, T1.word, T1.means, T0.count_marks, T0.last_mark FROM user_words_study T0 INNER JOIN words_list T1 ON T1.id = T0.word_id WHERE T0.user_id = ? ORDER BY RANDOM()", user.Id).QueryRows(&userWords)
	} else {
		o.Raw("SELECT T0.id, T1.word, T1.means, T0.count_marks, T0.last_mark FROM user_words_study T0 INNER JOIN words_list T1 ON T1.id = T0.word_id WHERE T0.user_id = ? ORDER BY T0.count_marks DESC, T1.id OFFSET ?", user.Id, const_conf.SyncLoadOffset).QueryRows(&userWords)
	}
	for _, v := range userWords {
		if v.CountMarks > 0 {
			tempWord = &const_conf.UserWordsJson{
				Id:         v.Id,
				Word:       v.Word,
				Means:      v.Means,
				CountMarks: v.CountMarks,
				LastMark:   timeago.Chinese.Format(v.LastMark),
			}
		} else {
			tempWord = &const_conf.UserWordsJson{
				Id:         v.Id,
				Word:       v.Word,
				Means:      v.Means,
				CountMarks: v.CountMarks,
				LastMark:   "",
			}
		}
		returnUserWords = append(returnUserWords, tempWord)
	}
	return returnUserWords
}

func FindUserWordByWordId(user *User, wordId int) (*UserWordsStudy, bool) {
	var userWord UserWordsStudy

	o := orm.NewOrm()
	err := o.QueryTable(userWord).Filter("UserId", user.Id).Filter("id", wordId).One(&userWord)
	return &userWord, err != orm.ErrNoRows
}

func IncrWordMark(UserWord *UserWordsStudy, user *User) {
	o := orm.NewOrm()
	UserWord.CountMarks = UserWord.CountMarks + 1
	o.Update(UserWord, "CountMarks", "LastMark")
	o.Insert(&UserLogs{User: user, Content: fmt.Sprintf(const_conf.LogMarkWordFmt, UserWord.Word.Id)})
}

func DeleteWord(user *User, wordId int) {
	var userWord UserWordsStudy

	o := orm.NewOrm()
	o.QueryTable(userWord).Filter("UserId", user.Id).Filter("id", wordId).Delete()
}

func CountOfMarkedWords(user *User) int64 {
	var userWord UserWordsStudy

	o := orm.NewOrm()
	count, _ := o.QueryTable(userWord).Filter("UserId", user.Id).Filter("CountMarks__gt", 0).Count()
	return count
}

func CountOfUserWords(user *User) int64 {
	var userWord UserWordsStudy

	o := orm.NewOrm()
	count, _ := o.QueryTable(userWord).Filter("UserId", user.Id).Count()
	return count
}
