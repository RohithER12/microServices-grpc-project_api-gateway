package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

type RemoveCartRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
	UserId    int64 `json:"userId"`
}

func RemoveCart(ctx *gin.Context, c pb.CartServiceClient) {
	body := RemoveCartRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.RemoveCart(context.Background(), &pb.RemoveCartRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    body.UserId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
