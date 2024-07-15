package main

import (
	"fmt"
	"gin_fw/src/database"
	"gin_fw/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)
	r := gin.Default()

	routes := route.NewRouteFacade()
	routes.SetupRoutes(r)

	database.Connect()

	appPort := os.Getenv("APP_PORT")
	address := fmt.Sprintf(":%s", appPort)
	r.Run(address)
}
