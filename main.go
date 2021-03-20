package main

import (
	"helpnow/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the routes
	r := routers.InitializeRoutes()

	// Start serving the application
	r.Run(":" + os.Getenv("PORT"))

}
