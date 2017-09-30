package const_conf

const (
	ServiceScheme = "https://"
	DomainName    = "gre3000.jolly23.com"

	HttpPort  = 8080
	HttpsPort = 8090

	IsEnableHTTP  = true
	IsEnableHTTPS = false

	SSLCertFile = "deploy_conf/certs/cert.pem"
	SSLKeyFile  = "deploy_conf/certs/cert.key"

	BeeGoViewsPath     = "controller_views"
	BeeGoConfiguration = "deploy_conf/beego_main.conf"
	BeeGoOrmAlias      = "default"
	BeeGoOrmMaxConn    = 30
	BeeGoOrmMaxIdle    = 15

	IsEnableXSRF    = true
	BeeGoXSRFKey    = "ASb&ADGaEmGQnp2XdTEc5NFw0Al0bKx$P1o61eJJF7$2b$1EoETzKXQuYh"
	BeeGoXSRFExpire = 3600



	RedisPort = "6379"
	RedisAddress = "link-docker-redis" + ":" + RedisPort
	RedisPassword = ""
	RedisNumber   = 0

	CookieSecure  = "$2m094FKSzyBj1DN27Ib$12$Fw0Al0bKGX9XuarHQzGDmtOSyeLWnfSbEc5N&AD"
	WebCookieName = "IsLogin"

	LogsMethod = "file"
	LogsConfig = `{"filename":"logs/site.log"}`


	PageSize = 20
)
