package main

import (
	"log"
	"os"
	"tes_project_siesta/configs"
	"tes_project_siesta/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configs.InitDB()

	r := routes.Routes(gin.Default())

	port := os.Getenv("PORT_SERVICES")
	r.Run(":" + port)
}
