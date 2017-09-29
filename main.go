package main

import (
	"GRE3000/const_conf"
	"GRE3000/models"
	_ "GRE3000/routers"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	beego.LoadAppConfig("ini", const_conf.BeeGoConfiguration)
	beego.SetLogger(const_conf.LogsMethod, const_conf.LogsConfig)

	dbSource := fmt.Sprintf("sslmode=%s host=%s port=%s dbname=%s user=%s password=%s",
		const_conf.PgSslMode,
		const_conf.PgHostAddress,
		const_conf.PgHostPort,
		const_conf.DatabaseName,
		const_conf.PgUserName,
		const_conf.PgPassword,
	)
	orm.RegisterDataBase(
		const_conf.BeeGoOrmAlias, const_conf.DatabaseType, dbSource,
		const_conf.BeeGoOrmMaxIdle, const_conf.BeeGoOrmMaxConn,
	)
	orm.RegisterModel(
		new(models.Common),
		new(models.WordsList),
		new(models.UsersList),
		new(models.UserLogs),
		new(models.UserWordsStudy),
	)
	orm.DefaultTimeLoc = time.UTC
	orm.RunSyncdb(const_conf.BeeGoOrmAlias, false, true)
	orm.Debug = true

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
}

func main() {
	beego.Run()
}
