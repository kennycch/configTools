package dao

import "config_tools/app/request"

type Game struct {
	Model
	Name       string `gorm:"type:varchar(50);not null;default:'';comment:游戏名" json:"name"`
	Background string `gorm:"type:varchar(255);not null;default:'';comment:游戏背景图" json:"background"`
	ClientGit  string `gorm:"type:varchar(255);not null;comment:前端git地址" json:"clientGit"`
	ServerGit  string `gorm:"type:varchar(255);not null;comment:后端git地址" json:"serverGit"`
	ExcelGit   string `gorm:"type:varchar(255);not null;comment:Excel git地址" json:"excelGit"`
}

func init() {
	tables = append(tables, createTable{
		table:   &Game{},
		comment: "游戏表",
	})
}

// GetGameList 获取账号列表
func GetGameList(req *request.ListBaseRequest) ([]*Game, int64) {
	list, count := []*Game{}, int64(0)
	quest := DB.Model(&Game{}).
		Select([]string{
			"id",
			"name",
			"background",
			"client_git",
			"server_git",
			"excel_git",
			"created_at",
		})
	quest.Count(&count)
	quest.Offset(GetOffset(req.Page, req.PageSize)).
		Limit(req.PageSize).
		Order("id desc").
		Find(&list)
	return list, count
}

// GetGameById 根据Id获取游戏
func GetGameById(game *Game) error {
	return DB.Where("id = ?", game.Id).
		First(game).
		Error
}

// GetGameByName 根据游戏名获取游戏
func GetGameByName(game *Game) error {
	return DB.Where("name = ?", game.Name).
		First(game).
		Error
}

// CreateGame 创建游戏
func CreateGame(game *Game) error {
	return DB.Create(game).
		Error
}

// GameUpdate 游戏更新
func GameUpdate(id uint32, updates map[string]interface{}) error {
	return DB.Model(&Game{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}
