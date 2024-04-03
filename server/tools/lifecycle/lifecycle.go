package lifecycle

import (
	"sort"
)

// 启动顺序，从高优先级到低优先级
func Start() {
	sort.SliceStable(lifecycles, func(i, j int) bool {
		return lifecycles[i].Priority() > lifecycles[j].Priority()
	})
	for _, lifecycle := range lifecycles {
		lifecycle.Start()
	}
}

// 停止顺序，跟启动顺序相反
func Stop() {
	sort.SliceStable(lifecycles, func(i, j int) bool {
		return lifecycles[i].Priority() < lifecycles[j].Priority()
	})
	for _, lifecycle := range lifecycles {
		lifecycle.Stop()
	}
}

// 注册服务
func AddLifecycle(lifecycle Lifecycle) {
	lifecycles = append(lifecycles, lifecycle)
}
