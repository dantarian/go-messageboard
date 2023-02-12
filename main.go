package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"pencethren/go-messageboard/config"
	"pencethren/go-messageboard/server"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(fmt.Errorf("failed to configure application: %w", err))
	}

	var db *sql.DB
	if config.Database.Type == "postgres" {
		db, err = sql.Open("postgres", dbConnectionString(config))
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}
	}

	s := server.NewServer(config, db)

	s.Run()
}

func dbConnectionString(c *config.Config) string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}
