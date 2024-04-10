package request

type GameCreateRequest struct {
	Name       string `form:"name" json:"name" binding:"required,max=50"`
	Background string `form:"background" json:"background" binding:"required,max=255"`
}

type GameEditRequest struct {
	GameCreateRequest
	DetailBaseRequest
}
