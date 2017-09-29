package models

type Common struct {
	Id           int    `orm:"pk;auto"`
	TheTableName string `orm:"size(256);column(table_name)"`
	ColumnName   string `orm:"size(256);"`
	ValueCode    string `orm:"size(256)"`
	CodeMean     string `orm:"size(256)"`
}

func (u *Common) TableName() string {
	return "db_common"
}
