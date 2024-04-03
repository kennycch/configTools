package lifecycle

// 执行优先级（从高到低）
const (
	HighPriority   = 10000000
	NormalPriority = 5000000
	LowPriority    = 0
)

type Lifecycle interface {
	Start()
	Priority() uint32
	Stop()
}

// 生命周期管理器
var (
	lifecycles = []Lifecycle{}
)
