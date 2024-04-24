package go_struct

type GoStruct struct {
	mainStruct   *Struct
	extraStructs []*Struct
	fileName     string
}

type Struct struct {
	lowName string
	upName  string
	keys    []*key
}

type key struct {
	annotation string // 注释
	typeStr    string // 类型
	lowName    string // 小驼峰
	upName     string // 大驼峰
	json       string // json字段
}

var (
	arrayOrder = []string{
		"[]bool",
		"[]int32",
		"[]float64",
		"[]string",
	}
)
