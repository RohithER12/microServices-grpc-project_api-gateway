package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func SortByPrice(ctx *gin.Context, c pb.ProductServiceClient) {
	sortParam := ctx.DefaultQuery("sort", "asc")

	sortAscending := sortParam == "asc"

	res, err := c.SortByPrice(context.Background(), &pb.SortByPriceRequest{
		Sort: sortAscending,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
