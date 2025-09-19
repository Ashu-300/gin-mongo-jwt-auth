package main

import (
	"gin-jwt-auth/database"
	"gin-jwt-auth/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	app := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}
	database.ConnectDb()
	routes.UserRouter(app)

	app.Run(":3000")
	
}