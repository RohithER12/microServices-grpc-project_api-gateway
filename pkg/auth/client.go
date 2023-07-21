package auth

import (
	"fmt"

	"github.com/RohithER12/api-gateway/pkg/auth/pb"
	"github.com/RohithER12/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
