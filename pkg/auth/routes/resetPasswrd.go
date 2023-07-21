package routes

import (
	"context"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type ResetPasswrdRequestBody struct {
	MobileNo string `json:"phone_number"`
}

type ResetPasswordValidationRequestBody struct {
	Key      string `json:"key"`
	Otp      string `json:"otp"`
	Password string `json:"password"`
}

func ResetPasswrd(ctx *gin.Context, c pb.AuthServiceClient) {
	body := ResetPasswrdRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ResetPassword(context.Background(), &pb.ResetPasswordRequest{

		PhoneNumber: body.MobileNo,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}

func ResetPasswordValidation(ctx *gin.Context, c pb.AuthServiceClient) {
	body := ResetPasswordValidationRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ResetPasswordValidation(context.Background(), &pb.ResetPasswordValidationRequest{
		Key:      body.Key,
		Otp:      body.Otp,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
