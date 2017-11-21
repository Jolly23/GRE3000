package utils

import (
	"GRE3000/models"
	"GRE3000/utils"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"github.com/xeonx/timeago"
	"time"
)

func FormatTime(time time.Time) string {
	return timeago.Chinese.Format(time)
}

func Markdown(content string) string {
	return string(blackfriday.MarkdownCommon([]byte(utils.NoHtml(content))))
}

func HasPermission(userId int, name string) bool {
	return models.FindPermissionByUserIdAndPermissionName(userId, name)
}

func init() {
	beego.AddFuncMap("FuncFormatTimeAgo", FormatTime)
	beego.AddFuncMap("FuncMarkDown", Markdown)
	beego.AddFuncMap("FuncHasPermission", HasPermission)
}
