package main

import (
	"fmt"
	"pencethren/go-messageboard/config"
	"pencethren/go-messageboard/server"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(fmt.Errorf("failed to configure application: %w", err))
	}

	s := server.NewServer(config)

	s.Run()
}
