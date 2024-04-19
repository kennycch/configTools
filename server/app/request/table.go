package request

type TableListRequest struct {
	ListBaseRequest
	GameId  uint32 `form:"gameId" json:"gameId"`
	Name    string `form:"name" json:"name"`
	Comment string `form:"comment" json:"comment"`
	Status  uint8  `form:"status" json:"status"`
}

type TableCreateRequest struct {
	GameId    uint32               `form:"gameId" json:"gameId" binding:"required,gte=1"`
	Name      string               `form:"name" json:"name" binding:"required,max=50"`
	Comment   string               `form:"comment" json:"comment" binding:"required,max=50"`
	TableType uint8                `form:"tableType" json:"tableType" binding:"required,gte=1,lte=2"`
	Fields    []CreateFieldRequest `form:"fields" json:"fields"`
}

type FieldRequest struct {
	Name      string `form:"name" json:"name" binding:"required,max=50"`
	Chinese   string `form:"chinese" json:"chinese" binding:"required,max=50"`
	Comment   string `form:"comment" json:"comment" binding:"max=255"`
	FieldType uint8  `form:"fieldType" json:"fieldType" binding:"required,gte=1,lte=9"`
	ParentId  int    `form:"parentId" json:"parentId"`
	Example   string `form:"example" json:"example"`
}

type CreateFieldRequest struct {
	FieldRequest
	Id int `form:"id" json:"id" binding:"required,lt=0"`
}
