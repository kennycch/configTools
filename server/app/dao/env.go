package dao

type Env struct {
	Model
	GameId   uint32 `gorm:"index:gameId;comment:游戏Id" json:"gameId"`
	EnvType  uint8  `gorm:"comment:环境类型：1：开发环境、2：预发布环境、3：正式环境" json:"envType"`
	JsonPath string `gorm:"type:varchar(255);not null;default:'';comment:前端git地址" json:"jsonPath"`
	GoPath   string `gorm:"type:varchar(255);not null;default:'';comment:后端git地址" json:"goPath"`
}

func init() {
	tables = append(tables, createTable{
		table:   &Env{},
		comment: "环境表",
	})
}
