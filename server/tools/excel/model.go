package excel

import (
	"config_tools/app/dao"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Local struct {
	cap int
	row int
}

type excel struct {
	table      *dao.Table
	fields     []*dao.Field
	file       *excelize.File
	titleStyle int
	cellStyle  int
}

var (
	// 列前缀
	caps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 字段类型映射
	typeMap = map[uint8]string{
		1: "bool",
		2: "int32",
		3: "float64",
		4: "string",
		5: "[]bool",
		6: "[]int32",
		7: "[]float64",
		8: "[]string",
		9: "[]object",
	}
	// 表类型映射
	tableTypeMap = map[uint8]string{
		1: "array",
		2: "object",
	}
	// 头单元格样式
	headStyle = `{
		"border":[
			{"type":"left","color":"000000", "style":1},
			{"type":"top","color":"000000", "style":1},
			{"type":"bottom","color":"000000", "style":1},
			{"type":"right","color":"000000", "style":1}
		],
		"font":{
			"bold":true,
			"size":14
		},
		"fill":{
			"type":"pattern",
			"pattern":1,
			"color":[
				"BDD7EE"
			]
		},
		"alignment":{
			"horizontal":"center",
			"vertical":"center",
			"justify_last_line":true,
			"wrap_text":true
		}
	}`
	// 标题单元格样式
	titleStyle = `{
		"border":[
			{"type":"left","color":"000000", "style":1},
			{"type":"top","color":"000000", "style":1},
			{"type":"bottom","color":"000000", "style":1},
			{"type":"right","color":"000000", "style":1}
		],
		"font":{
			"bold":true,
			"size":13
		},
		"fill":{
			"type":"pattern",
			"pattern":1,
			"color":[
				"FFFF00"
			]
		},
		"alignment":{
			"horizontal":"center",
			"vertical":"center",
			"justify_last_line":true,
			"wrap_text":true
		}
	}`
	// 值单元格样式
	cellStyle = `{
		"border":[
			{"type":"left","color":"000000", "style":1},
			{"type":"top","color":"000000", "style":1},
			{"type":"bottom","color":"000000", "style":1},
			{"type":"right","color":"000000", "style":1}
		],
		"font":{
			"size":12
		},
		"alignment":{
			"horizontal":"center",
			"vertical":"center",
			"justify_last_line":true,
			"wrap_text":true
		}
	}`
)
