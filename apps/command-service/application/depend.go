package application

import (
	"github.com/tortillaproduction/go-microservices/apps/command-service/application/impl"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler"
	"go.uber.org/fx"
)

var SrvDepend = fx.Options(
	sqlboiler.RepDepend,
	fx.Provide(
		impl.NewCategoryServiceImpl,
		impl.NewProductServiceImpl,
	),
)
