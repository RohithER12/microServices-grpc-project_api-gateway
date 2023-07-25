package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

type AddCartRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
	UserId    int64 `json:"userId"`
}

func AddCart(ctx *gin.Context, c pb.CartServiceClient) {
	body := AddCartRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddCart(context.Background(), &pb.AddCartRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    body.UserId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
