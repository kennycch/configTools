package dao

import "config_tools/app/request"

type Table struct {
	Model
	GameId    uint32 `gorm:"index:gameId;comment:游戏Id" json:"gameId"`
	Name      string `gorm:"type:varchar(50);not null;comment:表名" json:"name"`
	Comment   string `gorm:"type:varchar(50);not null;comment:表注释" json:"comment"`
	Hash      string `gorm:"type:varchar(50);not null;comment:哈希签名" json:"hash"`
	TableType uint8  `gorm:"comment:表类型：1：数组、2：对象" json:"tableType"`
	Status    uint8  `gorm:"comment:表状态：1：已生成，2：已上传，3：已发布" json:"status"`
	Version   uint32 `gorm:"comment:版本号" json:"version"`
}

type TableList struct {
	Table
	GameName string `gorm:"column:gameName" json:"gameName"`
}

func init() {
	tables = append(tables, createTable{
		table:   &Table{},
		comment: "配置表",
	})
}

func GetTableList(req *request.TableListRequest) ([]*TableList, int64) {
	list, count := []*TableList{}, int64(0)
	quest := DB.Model(&Table{}).
		Select([]string{
			"table.id",
			"table.name",
			"table.comment",
			"table.hash",
			"table.version",
			"table.table_type",
			"table.status",
			"game.name as gameName",
		}).
		Joins("join game ON game.id = table.game_id")
	if req.GameId > 0 {
		quest.Where("game.id = ?", req.GameId)
	}
	if req.Name != "" {
		quest.Where("table.name like ?", "%"+req.Name+"%")
	}
	if req.Name != "" {
		quest.Where("table.comment = ?", "%"+req.Comment+"%")
	}
	if req.Status > 0 {
		quest.Where("table.status = ?", req.Status)
	}
	quest.Count(&count)
	quest.Offset(GetOffset(req.Page, req.PageSize)).
		Limit(req.PageSize).
		Order("table.id desc").
		Find(&list)
	return list, count
}

func GetTableByGameIdAndName(table *Table) error {
	return DB.Where("game_id = ? AND name = ?", table.GameId, table.Name).
		First(table).
		Error
}
