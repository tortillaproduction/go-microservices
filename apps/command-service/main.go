package main

import (
	"time"

	"github.com/tortillaproduction/go-microservices/apps/command-service/presen"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		presen.CommandDepend,
		fx.StartTimeout(60*time.Second),
	).Run()
}
