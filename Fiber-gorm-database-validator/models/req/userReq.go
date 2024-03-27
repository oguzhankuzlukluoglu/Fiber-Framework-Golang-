package req

type UserReq struct {
	Name  string `json:"name validate:"required"`
	Email string `json:"email validate:"required,email,min=3,max=32"`
}
