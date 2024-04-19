package config

import (
	"github.com/go-ini/ini"
)

// 配置赋值
func MapConfig(env *ini.File) {
	env.Section("app").MapTo(App)
	env.Section("log").MapTo(Log)
	env.Section("http").MapTo(Http)
	env.Section("sign").MapTo(Sign)
	env.Section("jwt").MapTo(Jwt)
	env.Section("redis").MapTo(Redis)
	env.Section("mysql").MapTo(Mysql)
	env.Section("git").MapTo(Git)
}
