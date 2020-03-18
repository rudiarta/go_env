package main

import (
	"github.com/rudirahardian/go_env/app/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	dotenv := goDotEnvVariable("PORT")
	routes.RouteInit(dotenv)
}