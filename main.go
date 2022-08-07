package main

import (
	"pencethren/go-messageboard/config"
	"pencethren/go-messageboard/server"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(config.Module, server.Module)
	app.Run()
}
