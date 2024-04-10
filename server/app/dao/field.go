package dao

type Field struct {
	Model
	TableId   uint32 `gorm:"index:tableId;comment:配置表Id" json:"tableId"`
	Name      string `gorm:"type:varchar(50);comment:字段名" json:"name"`
	Chinese   string `gorm:"type:varchar(50);comment:字段中文" json:"chinese"`
	Comment   string `gorm:"type:varchar(255);comment:字段注释" json:"comment"`
	FieldType uint8  `gorm:"comment:字段类型：1：布尔，2：整型，3：浮点型，4：字符串，5：数组布尔，6：数组整形，7：数组浮点型，8：数组字符串，9：数组对象" json:"fieldType"`
	ParentId  uint32 `gorm:"index:parentId;comment:上级Id" json:"parentId"`
	Example   string `gorm:"comment:示例" json:"example"`
}

func init() {
	tables = append(tables, createTable{
		table:   &Field{},
		comment: "字段表",
	})
}
