package main

import (
	"config_tools/app/dao"
	"config_tools/config"
	"config_tools/tools/git"
	"config_tools/tools/lifecycle"
	"config_tools/tools/log"
	"config_tools/tools/net"
	"config_tools/tools/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 服务注册器
	register()
	// 服务开启事件
	lifecycle.Start()
	// 服务结束事件
	defer lifecycle.Stop()
	// 信号监听
	loop()
}

// 服务注册器
func register() {
	// 加载配置
	lifecycle.AddLifecycle(config.NewConfigRegister())
	// 初始化日志
	lifecycle.AddLifecycle(log.NewLogRegister())
	// 开启服务
	lifecycle.AddLifecycle(net.NewNetRegister())
	// 初始化Redis
	lifecycle.AddLifecycle(redis.NewRedisRegister())
	// 初始化MySql
	lifecycle.AddLifecycle(dao.NewMysqlRegister())
	// 克隆前后端项目
	lifecycle.AddLifecycle(git.NewGitRegister())
}

// 信号监听
func loop() {
	log.Info("server started")
	signals := make(chan os.Signal, 1)
	// kill -9 无法被捕获
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
label:
	for {
		s := <-signals
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			log.Info("server close... ...")
			break label
		case syscall.SIGHUP:
			continue
		}
	}
}
