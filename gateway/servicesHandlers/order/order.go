package order

type OrderReq struct {
	Id       int `form:"id" json:"id"`
	Quantity int `form:"quantity" json:"quantity" binding:"required"`
	UserId   int `form:"userId" json:"userId" binding:"required"`
}
