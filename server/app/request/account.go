package request

type LoginRequest struct {
	Account  string `form:"account" json:"account" binding:"required,min=4,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
	Captcha  string `form:"captcha" json:"captcha" binding:"required,min=6,max=6"`
	Code     string `form:"code" json:"code" binding:"required,min=10,max=30"`
}

type ChangePassword struct {
	OldPassword     string `form:"oldPassword" json:"oldPassword" binding:"required,min=6,max=30"`
	NewPassword     string `form:"newPassword" json:"newPassword" binding:"required,min=6,max=30"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required,min=6,max=30"`
}

type AccountListRequest struct {
	ListBaseRequest
	Account string `form:"account" json:"account"`
	Status  *uint8 `form:"status" json:"status"`
}

type AccountCreateRequest struct {
	Account  string `form:"account" json:"account" binding:"required,min=4,max=30"`
	IsSupper uint8  `form:"isSupper" json:"isSupper" binding:"required,lte=1"`
}
