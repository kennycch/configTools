package config

import "flag"

var (
	// 配置路径
	filePath = flag.String("c", "./env.ini", "config file")

	// 总配置
	App = &struct {
		AppName string // 进程名
		Debug   bool   // 是否开启debug模式
	}{}

	// Http配置
	Http = &struct {
		Port int32 // 服务监听端口
	}{}

	// 日志配置
	Log = &struct {
		LogPath string // 输出日志路径
		LogDay  int    // 日志存放天数
	}{}

	// 签名配置
	Sign = &struct {
		SignKey string // 签名秘钥
		TimeOut int64  // 签名超时时间（秒）
	}{}

	// Jwt配置
	Jwt = &struct {
		SecretKey string // token解析Key
		TimeOut   int64  // token有效时间（小时）
		Issuer    string // 发行者
	}{}

	// Redis配置
	Redis = &struct {
		Addr     string // Redis地址
		PassWord string // 密码
		DB       int32  // 选择库
	}{}

	// Mysql配置
	Mysql = &struct {
		Addr     string // Mysql地址
		Account  string // 账号
		Password string // 密码
		DataBase string // 库
	}{}

	// Git配置
	Git = &struct {
		Name    string // Git用户名
		Email   string // Git邮箱
		SshPath string // SSH秘钥路径
	}{}
)
