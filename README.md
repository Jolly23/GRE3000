# GRE3000

`没什么意义的README`

### 用docker 或者 docker-compose运行，我都写好了，直接能用

> 采用了4容器负载均衡，为的是实验多机部署。

使用前请创建 const_conf/private.go 并加入PostgreSql数据库地址

```Go
package const_conf

const (
	ServerIP = "localhost"

	DatabaseType  = "postgres"
	PgSslMode     = "disable"
	PgHostAddress = "111.111.111.111"
	PgHostPort    = "5432"
	DatabaseName  = "GRE3000"
	PgUserName    = "JOLLY"
	PgPassword    = "HW(D*H(AWD"
)

```

> 可以更换为其他数据库，需更改main.go文件，具体请参考[Beego框架ORM](https://beego.me/docs/mvc/model/overview.md)

### 运行

```bash
git clone https://github.com/Jolly23/GRE3000.git
cd GRE3000
vim const_conf/private.go
docker-compose build
docker-compose up -d
```

`可根据写好的nginx配置文件来引导访问`
`数据库自动建表完毕后，要记得导入词表，在deploy_conf文件夹中，然后执行一下pg-gre300.sql中的所有语句即可`

### TODO
> 1. 单词表页面一次性发送整个GRE词表，大小大约在3M，用户体验不太好，下一步用js异步加载词表，但完成3000个div加载性能同样较差，正在寻求解决方案。（如果服务器带宽大一点的话这个可能就不是问题了）

> 2. 为用户加入全部显示中文意思的词表按钮

> 3. 增强管理员视图

> 4. 使用kubernetes替换docker-compose