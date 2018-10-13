# [GRE3000](https://gre.jolly23.com/)

`很神圣的README`

## 用docker 或者 docker-compose运行，我都写好了，直接能用

> 采用了4容器负载均衡，为的是实验多机部署。

代码clone后编译前请创建 const_conf/private.go 并加入PostgreSQL数据库地址

```Go
package const_conf

const (
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


## 运行

```bash
git clone https://github.com/Jolly23/GRE3000.git
cd GRE3000
vim const_conf/private.go   # 写入上述go文件，配置数据库连接信息
docker-compose build
docker-compose up -d
```

`可根据写好的Nginx配置文件来引导访问`
`数据库自动建表完毕后，要记得导入词表，在deploy_conf文件夹中，然后执行一下pg-gre300.sql中的所有语句即可`


## TODO LIST

### 1. 单词表页面一次性发送整个GRE词表，大小大约在3M，用户体验不太好，下一步用js异步加载词表，但完成3000个div加载性能同样较差，正在寻求解决方案。（如果服务器带宽大一点的话这个可能就不是问题了）
> 已解决
> 默认加载30个单词的HTML，然后通过js请求json加载剩余单词，数据瘦身80%
> 通过html页面头部替代尾部提前加载js函数，使用户可以在网页加载时使用可见单词的全部按钮功能

### 2. 为用户加入全部显示中文意思的词表按钮
> 已完成

### 3. 增强管理员视图

### 4. 使用Kubernetes替换docker-compose


## License
The MIT License(https://opensource.org/licenses/MIT)

请自由地享受和参与开源

## 贡献

如果你有好的意见或建议，欢迎提issue或pull request
