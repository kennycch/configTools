package request

type TableListRequest struct {
	ListBaseRequest
	GameId  uint32 `form:"gameId" json:"gameId"`
	Name    string `form:"name" json:"name"`
	Comment string `form:"comment" json:"comment"`
	Status  uint8  `form:"status" json:"status"`
}

