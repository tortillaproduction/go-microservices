package main

import (
	"github.com/tortillaproduction/go-microservices/apps/command-service/presen"
	"go.uber.org/fx"
)

func main() {
	fx.New(presen.CommandDepend).Run()
}
