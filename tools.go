//go:build tools

package main

import (
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"
)
