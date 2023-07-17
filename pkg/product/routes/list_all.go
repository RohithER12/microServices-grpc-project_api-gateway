package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func ListProducts(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.ListProducts(context.Background(), &pb.ListProductsRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
