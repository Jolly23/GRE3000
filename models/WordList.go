package models

import (
	"github.com/Jolly23/GRE3000/const_conf"
	"github.com/astaxie/beego/orm"
)

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
	o.QueryTable(tableOfWords).OrderBy("id").Limit(const_conf.SyncLoadOffset).All(&allWords)
	return allWords
}

func LoadRawWordsJson(random bool) *[]orm.Params {
	o := orm.NewOrm()
	var maps []orm.Params
	if random {
		o.Raw("SELECT word, means FROM words_list ORDER BY RANDOM()").Values(&maps)
	} else {
		o.Raw("SELECT word, means FROM words_list ORDER BY id OFFSET ?", const_conf.SyncLoadOffset).Values(&maps)
	}
	return &maps
}

func FindWordById(id int) *WordsList {
	o := orm.NewOrm()
	var word WordsList
	o.QueryTable(word).Filter("Id", id).One(&word)
	return &word
}
