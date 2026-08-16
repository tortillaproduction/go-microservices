package prepare

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

func CommandServiceLifecycle(lifecycle fx.Lifecycle, server *CommandServer) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := handler.DBConnect(); err != nil {
				panic(err)
			}
			port := 8082
			listner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				return err
			}
			reflection.Register(server.Server)

			go func() {
				log.Printf("Command Server started on port: %v", port)
				server.Server.Serve(listner)
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			server.Server.GracefulStop()
			log.Println("Command Server stopped")
			return nil
		},
	})
}
