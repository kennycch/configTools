package redis

import (
	"config_tools/config"

	"github.com/go-redis/redis"
)

func RedisInit() {
	con := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,     // 连接地址（含端口）
		Password: config.Redis.PassWord, // 密码
		DB:       int(config.Redis.DB),  // 选择的库
		PoolSize: 30,                    // 连接池数量
	})
	_, err := con.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	RD = con
}
