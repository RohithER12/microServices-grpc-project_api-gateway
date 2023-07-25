package product

import (
	"fmt"

	"github.com/RohithER12/api-gateway/pkg/auth"
	"github.com/RohithER12/api-gateway/pkg/config"
	"github.com/RohithER12/api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
	routes.GET("/products", svc.ListProducts)
	routes.GET("/search", svc.Search)
	routes.GET("/sortByPrice", svc.SortByPrice)

}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	fmt.Println("a search\t\thrrrre\t\tString\n\n\t\t")

	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) ListProducts(ctx *gin.Context) {
	routes.ListProducts(ctx, svc.Client)
}

func (svc *ServiceClient) Search(ctx *gin.Context) {
	fmt.Println("a searchString\n\n\t\t")

	routes.Search(ctx, svc.Client)
}

func (svc *ServiceClient) SortByPrice(ctx *gin.Context) {
	routes.SortByPrice(ctx, svc.Client)
}
