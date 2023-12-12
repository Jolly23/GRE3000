package main

import (
	"github.com/gofiber/template/html/v2"
	"github.com/russross/blackfriday"
	"github.com/xeonx/timeago"
	"html/template"
	"strings"
	"time"
)

func RenderOS() *html.Engine {

	// Create a new engine by passing the template folder
	// and template extension using <engine>.New(dir, ext string)
	engine := html.New("./views", ".tpl")

	//// We also support the http.FileSystem interface
	//// See examples below to load templates from embedded files
	//engine := html.NewFileSystem(http.Dir("./views"), ".html")

	// Reload the templates on each render, good for development
	//engine.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	//engine.Debug(true) // Optional. Default: false

	// Layout defines the variable name that is used to yield templates within layouts
	engine.Layout("embed") // Optional. Default: "embed"

	// Delims sets the action delimiters to the specified strings
	engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// AddFunc adds a function to the template's global function map.
	engine.AddFunc("FuncFormatTimeAgo", formatTime)
	engine.AddFunc("FuncMarkDown", markdown)
	//engine.AddFunc("FuncHasPermission", hasPermission)
	engine.AddFunc("str2html", str2html)

	return engine
}

func formatTime(time time.Time) string {
	return timeago.Chinese.Format(time)
}

func markdown(content string) string {
	return string(blackfriday.MarkdownCommon([]byte(noHtml(content))))
}

//func hasPermission(userId int, name string) bool {
//	return true
//	//return models.FindPermissionByUserIdAndPermissionName(userId, name)
//}

func str2html(raw string) template.HTML {
	return template.HTML(raw)
}

func noHtml(str string) string {
	return strings.Replace(strings.Replace(str, "<script", "&lt;script", -1), "script>", "script&gt;", -1)
}
