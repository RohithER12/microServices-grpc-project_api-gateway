package cart

import (
	"fmt"

	"github.com/RohithER12/api-gateway/pkg/cart/pb"
	"github.com/RohithER12/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CartServiceClient
}

func InitServiceClient(c *config.Config) pb.CartServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CartSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCartServiceClient(cc)
}
