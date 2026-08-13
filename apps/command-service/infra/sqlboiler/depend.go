package sqlboiler

import (
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/repository"
	"go.uber.org/fx"
)

var RepDepend = fx.Options(
	fx.Provide(
		repository.NewCategoryRepositorySQLBoiler(),
		repository.NewProductRepositorySQLBoiler(),
	),
)
