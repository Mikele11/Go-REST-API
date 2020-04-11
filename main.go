package main

import (
	"log"
	"os"

	"./routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	var appPort = goDotEnvVariable("PORT")

	if appPort == "" {
		appPort = "1111"
	}

	routes.SetupServer(appPort)
}
