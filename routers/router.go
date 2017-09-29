package routers

import (
	"GRE3000/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/", &controllers.WordsController{}, "get:Index")
	beego.Router("/test", &controllers.TestController{}, "get:Index")

	beego.Router("/cookie", &controllers.CookieController{}, "get:Test")
	beego.Router("/cookie2", &controllers.CookieController{}, "get:Test2")

}
