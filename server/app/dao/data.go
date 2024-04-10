package dao

type Data struct {
	Id      uint64 `gorm:"autoIncrement" json:"id"`
	FieldId uint32 `gorm:"index:fieldId;comment:字段表Id" json:"fieldId"`
	EnvId   uint32 `gorm:"index:envId;comment:环境Id" json:"envId"`
	Value   string `gorm:"comment:值" json:"value"`
}

func init() {
	tables = append(tables, createTable{
		table:   &Data{},
		comment: "数值表",
	})
}

// 数据填充
func (d *Data) dataFilling() {

}
