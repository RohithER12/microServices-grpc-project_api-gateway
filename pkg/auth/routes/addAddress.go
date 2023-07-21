package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AddAddressRequestBody struct {
	DoorNo     string `json:"doorno"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
}

func AddAdress(ctx *gin.Context, c pb.AuthServiceClient) {
	body := AddAddressRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId, _ := ctx.Get("userId")

	res, err := c.AddAddress(context.Background(), &pb.AddAddressRequest{
		DoorNo:      body.DoorNo,
		City:        body.City,
		PostalCodev: body.PostalCode,
		State:       body.State,
		UserId:      userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
