package main

import (
	"os"

	"./routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	var appPort = os.Getenv("appPort")

	if appPort == "" {
		appPort = "1111"
	}

	routes.SetupServer(appPort)
}
