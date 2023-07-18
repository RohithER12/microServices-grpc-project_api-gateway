package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	MobileNo string `json:"phone_number"`
}

type RegisterValidateRequestBody struct {
	Key string `json:"key"`
	Otp string `json:"otp"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:       body.Email,
		Password:    body.Password,
		PhoneNumber: body.MobileNo,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}

func RegisterOTPValidation(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterValidateRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.RegisterOTPValidation(context.Background(), &pb.RegisterOTPValidationRequest{
		Key: body.Key,
		Otp: body.Otp,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
