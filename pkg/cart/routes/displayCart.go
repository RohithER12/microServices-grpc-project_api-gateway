package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

type DisplayCartRequestBody struct {
	UserId int64 `json:"userId"`
}

func DisplayCart(ctx *gin.Context, c pb.CartServiceClient) {
	res, err := c.DisplayCart(context.Background(), &pb.DisplayCartRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
