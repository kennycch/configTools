package redis

import "config_tools/tools/lifecycle"

type RedisRegister struct{}

func (r *RedisRegister) Start() {
	RedisInit()
}

func (r *RedisRegister) Priority() uint32 {
	return lifecycle.NormalPriority + 10000
}

func (r *RedisRegister) Stop() {
	RD.Close()
}

func NewRedisRegister() *RedisRegister {
	return &RedisRegister{}
}
