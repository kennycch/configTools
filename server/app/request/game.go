package request

type GameCreateRequest struct {
	Name       string `form:"name" json:"name" binding:"required,max=50"`
	Background string `form:"background" json:"background" binding:"required,max=255"`
	ClientGit  string `form:"clientGit" json:"clientGit" binding:"required,max=255"`
	ServerGit  string `form:"serverGit" json:"serverGit" binding:"required,max=255"`
	ExcelGit   string `form:"excelGit" json:"excelGit" binding:"required,max=255"`
}

type GameEditRequest struct {
	GameCreateRequest
	DetailBaseRequest
}
