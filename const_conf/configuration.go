package const_conf

import (
	"fmt"
)

const (
	ServiceScheme = "https://"
	//DomainName    = "gre.jolly23.com"
	DomainName = ""

	HttpPort  = 8080
	HttpsPort = 8090

	IsEnableHTTP  = true
	IsEnableHTTPS = false

	SSLCertFile = "deploy_conf/certs/cert.pem"
	SSLKeyFile  = "deploy_conf/certs/cert.key"

	BeeGoViewsPath     = "controller_views"
	BeeGoConfiguration = "deploy_conf/beego_main.conf"
	BeeGoOrmAlias      = "default"
	BeeGoOrmMaxConn    = 20
	BeeGoOrmMaxIdle    = 10

	IsEnableXSRF    = false
	BeeGoXSRFKey    = "ASb&ADGaEmGQnp2XdTEc5NFw0Al0bKx$P1o61eJJF7$2b$1EoETzKXQuYh"
	BeeGoXSRFExpire = 3600

	RedisPort     = "6379"
	RedisAddress  = "link-docker-redis" + ":" + RedisPort
	RedisPassword = ""
	RedisNumber   = 0

	CookieSecure  = "$2m094FKSzyBj1DN27Ib$12$Fw0Al0bKGX9XuarHQzGDmtOSyeLWnfSbEc5N&AD"
	WebCookieName = "IsLogin"

	LogsMethod = "file"
	LogsConfig = `{"filename":"logs/site.log"}`

	PageSize = 20

	MarkWordTimeLimit = 10
)

var DbSource string = fmt.Sprintf("sslmode=%s host=%s port=%s dbname=%s user=%s password=%s",
	PgSslMode,
	PgHostAddress,
	PgHostPort,
	DatabaseName,
	PgUserName,
	PgPassword,
)
