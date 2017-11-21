package main

import (
	"GRE3000/const_conf"
	"GRE3000/models"
	_ "GRE3000/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	_ "GRE3000/base/cache"
	_ "GRE3000/templates"
	_ "GRE3000/utils"
)

func init() {
	beego.LoadAppConfig("ini", const_conf.BeeGoConfiguration)
	beego.SetLogger(const_conf.LogsMethod, const_conf.LogsConfig)
	orm.RegisterDataBase(
		const_conf.BeeGoOrmAlias, const_conf.DatabaseType, const_conf.DbSource,
		const_conf.BeeGoOrmMaxIdle, const_conf.BeeGoOrmMaxConn,
	)
	orm.RegisterModel(
		new(models.Common),
		new(models.WordsList),
		new(models.UserLogs),
		new(models.UserWordsStudy),
		new(models.User),
		new(models.Topic),
		new(models.Section),
		new(models.Reply),
		new(models.ReplyUpLog),
		new(models.Role),
		new(models.Permission),
	)

	orm.DefaultTimeLoc = time.UTC
	orm.RunSyncdb(const_conf.BeeGoOrmAlias, false, true)
	orm.Debug = false

	beego.BConfig.Listen.EnableHTTP = const_conf.IsEnableHTTP
	beego.BConfig.Listen.HTTPPort = const_conf.HttpPort

	if const_conf.IsEnableHTTPS {
		beego.BConfig.Listen.EnableHTTPS = const_conf.IsEnableHTTPS
		beego.BConfig.Listen.HTTPSPort = const_conf.HttpsPort
		beego.BConfig.Listen.HTTPSCertFile = const_conf.SSLCertFile
		beego.BConfig.Listen.HTTPSKeyFile = const_conf.SSLKeyFile
	}

	if const_conf.IsEnableXSRF {
		beego.BConfig.WebConfig.EnableXSRF = const_conf.IsEnableXSRF
		beego.BConfig.WebConfig.XSRFKey = const_conf.BeeGoXSRFKey
		beego.BConfig.WebConfig.XSRFExpire = const_conf.BeeGoXSRFExpire
	}

	beego.BConfig.WebConfig.ViewsPath = const_conf.BeeGoViewsPath

	models.Init()
}

func main() {
	beego.Run()
}
