package cart

import (
	"github.com/RohithER12/api-gateway/pkg/auth"
	"github.com/RohithER12/api-gateway/pkg/cart/routes"
	"github.com/RohithER12/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/cart")
	routes.Use(a.AuthRequired)
	routes.POST("/addCart", svc.AddCart)
	routes.PATCH("/removeCart", svc.RemoveCart)
	routes.GET("displayCart", svc.DisplayCart)

}

func (svc *ServiceClient) AddCart(ctx *gin.Context) {
	routes.AddCart(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveCart(ctx *gin.Context) {
	routes.RemoveCart(ctx, svc.Client)
}

func (svc *ServiceClient) DisplayCart(ctx *gin.Context) {
	routes.DisplayCart(ctx, svc.Client)
}
