package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/o77tsen/den/initializers"
	"github.com/o77tsen/den/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect()
}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	routes.MessageRoutes(router)
	router.Run()
}
