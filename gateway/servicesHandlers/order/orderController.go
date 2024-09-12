package order

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IOrderController interface {
	CreateOrder(ctx *gin.Context)
}
type OrderController struct {
	orderService IOrderService
}

func NewOrderController(iOrderService IOrderService) *OrderController {
	return &OrderController{orderService: iOrderService}
}
func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var orderReq OrderReq
	err := ctx.Bind(&orderReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create order:%v", err.Error())})
		return
	}
	creationErr := o.orderService.CreateOrder(orderReq)
	if creationErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create order:%v", creationErr.Error())})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
