# GRE3000

`没什么意义的README`

### 用docker 或者 docker-compose运行，我都写好了，直接能用

使用前请创建 const_conf/private.go 并加入PostgreSql数据库地址


> 可以更换为其他数据库，具体请参考Beego(框架)[https://beego.me/docs/mvc/model/overview.md]
```Go
package const_conf

const (
	ServerIP = "60.205.209.112"

	DatabaseType  = "postgres"
	PgSslMode     = "disable"
	PgHostAddress = ServerIP
	PgHostPort    = "5432"
	DatabaseName  = "GRE3000"
	PgUserName    = "JOLLY"
	PgPassword    = "200800"
)

```
