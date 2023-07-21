package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

type UserProfileResponse struct {
	Email        string
	Phone_number string
}

type UserProfileRequest struct {
	UserId string
}

func UserProfile(ctx *gin.Context, c pb.AuthServiceClient) {
	userId, _ := ctx.Get("userId")

	res, err := c.UserProfile(context.Background(), &pb.UserProfileRequest{
		UserId: userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
