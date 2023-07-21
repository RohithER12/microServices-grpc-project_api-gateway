package auth

import (
	"github.com/RohithER12/api-gateway/pkg/auth/routes"
	"github.com/RohithER12/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	a := InitAuthMiddleware(svc)

	authRoutes := r.Group("/authRoutes")
	authRoutes.Use(a.AuthRequired)
	authRoutes.GET("/userProfile", svc.UserProfile)
	authRoutes.POST("/addAddress", svc.AddAddress)

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/registerValidation", svc.RegisterOTPValidation)
	routes.POST("/login", svc.Login)
	routes.POST("/resetPassword", svc.ResetPassword)
	routes.POST("/resetPasswordValidation", svc.ResetPasswordValidation)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) RegisterOTPValidation(ctx *gin.Context) {
	routes.RegisterOTPValidation(ctx, svc.Client)
}

func (svc *ServiceClient) ResetPassword(ctx *gin.Context) {
	routes.ResetPasswrd(ctx, svc.Client)
}

func (svc *ServiceClient) ResetPasswordValidation(ctx *gin.Context) {
	routes.ResetPasswordValidation(ctx, svc.Client)
}

func (svc *ServiceClient) UserProfile(ctx *gin.Context) {
	routes.UserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) AddAddress(ctx *gin.Context) {
	routes.AddAdress(ctx, svc.Client)
}
