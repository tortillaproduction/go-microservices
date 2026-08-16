package prepare

import (
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"google.golang.org/grpc"
)

type CommandServer struct {
	Server *grpc.Server
}

func NewCommandServer(category pbv1.CategoryCommandServer, product pbv1.ProductCommandServer) *CommandServer {
	server := grpc.NewServer()

	pbv1.RegisterCategoryCommandServer(server, category)
	pbv1.RegisterProductCommandServer(server, product)

	return &CommandServer{
		Server: server,
	}
}
