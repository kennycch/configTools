package request

type ListBaseRequest struct {
	Page     int `form:"page" json:"page" binding:"required,gte=1"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required,gte=10"`
}

type DetailBaseRequest struct {
	Id uint32 `form:"id" json:"id" binding:"required,gte=1"`
}
