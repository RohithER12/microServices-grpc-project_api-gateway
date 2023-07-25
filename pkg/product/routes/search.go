package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context, c pb.ProductServiceClient) {
	searchString := ctx.Query("search")
	fmt.Println("a searchString\n\n\t\t", searchString)

	res, err := c.Search(context.Background(), &pb.SearchRequest{
		Search: searchString,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}
