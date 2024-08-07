package order

type OrderReq struct {
	ItemId   int `form:"itemId" json:"itemId" binding:"required"`
	Quantity int `form:"quantity" json:"quantity" binding:"required"`
	UserId   int `form:"userId" json:"userId" binding:"required"`
}
