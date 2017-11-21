package cache

import (
	"GRE3000/const_conf"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	Redis cache.Cache
)

func init() {
	redisConf := fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%d","password":"%s"}`,
		const_conf.DomainName,
		const_conf.RedisAddress,
		const_conf.RedisNumber,
		const_conf.RedisPassword,
	)
	Redis, _ = cache.NewCache("redis", redisConf)
}
