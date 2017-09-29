package models

import "github.com/astaxie/beego/orm"

type WordsList struct {
	Id    int    `orm:"pk;auto;index"`
	Words string `orm:"size(64);index"`
	Means string `orm:"size(512);"`
}

func (u *Common) WordsList() string {
	return "words_list"
}

func LoadWords() []*WordsList {
	o := orm.NewOrm()
	var tableOfWords WordsList
	var allWords []*WordsList
	o.QueryTable(tableOfWords).Limit(-1).All(&allWords)
	return allWords
}
