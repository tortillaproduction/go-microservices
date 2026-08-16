package presen

import (
	"github.com/tortillaproduction/go-microservices/apps/command-service/application"
	"github.com/tortillaproduction/go-microservices/apps/command-service/presen/adapter"
	"github.com/tortillaproduction/go-microservices/apps/command-service/presen/prepare"
	"github.com/tortillaproduction/go-microservices/apps/command-service/presen/server"
	"go.uber.org/fx"
)

var CommandDepend = fx.Options(
	application.SrvDepend,
	fx.Provide(
		adapter.NewCategoryAdapterImpl,
		adapter.NewProductAdapterImpl,
		server.NewCategoryServer,
		server.NewProductServer,
		prepare.NewCommandServer,
	),
	fx.Invoke(prepare.CommandServiceLifecycle),
)
