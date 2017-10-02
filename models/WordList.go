package models

import "github.com/astaxie/beego/orm"

type WordsList struct {
	Id            int               `orm:"pk;auto;index"`
	Word          string            `orm:"size(64);index"`
	Means         string            `orm:"size(512);"`
	LinkUserStudy []*UserWordsStudy `orm:"reverse(many)"`
}

func (u *Common) WordsList() string {
	return "words_list"
}

func LoadRawWords() []*WordsList {
	o := orm.NewOrm()
	var tableOfWords WordsList
	var allWords []*WordsList
	o.QueryTable(tableOfWords).Limit(-1).All(&allWords)
	return allWords
}

func FindWordById(id int) *WordsList {
	o := orm.NewOrm()
	var word WordsList
	o.QueryTable(word).Filter("Id", id).One(&word)
	return &word
}
