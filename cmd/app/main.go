package main

import (
	"dbUtils/configs"
	"dbUtils/internal"
	"flag"
	"log"
)

var (
	port     = flag.Uint("port", 1433, "the database port")
	server   = flag.String("server", "hostname", "the database server")
	user     = flag.String("user", "tezis-test-service-user", "the database user")
	password = flag.String("password", "wrong_password", "the database password")
	database = flag.String("database", "database", "the database name")
	debug    = flag.Bool("debug", false, "enable debugging")
)

func main() {
	flag.Parse()

	cfg := configs.DatabaseConfiguration{
		*server, uint16(*port), *user, *password, *database,
	}

	appName := "Database Utils"
	log.Printf("Привет, %s!", appName)

	internal.RunApp(cfg.CreateConnStr())

}
