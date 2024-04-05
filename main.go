package main

import (
	"fmt"
	"net/http"
	"os"

	"publisher/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Jai Guru Dev")

	envErr := godotenv.Load()
	_ = envErr

	apiServer := gin.Default()

	//Cors setting
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"token", "content-type", "content-length", "accept-encoding", "cache-control", "payload"}
	config.AllowMethods = []string{"GET", "PUT", "POST", "OPTIONS", "PATCH", "HEAD"}
	// apiServer.Use(cors.New(config))

	//Routes grouping
	apiV1Server := apiServer.Group("/api/v1")
	apiV1Server.Use(cors.New(config))

	apiV1Server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "v1.0 running !!"})
	})

	apphandler := handler.NewAppHandler()

	apiV1Server.POST("/produce-message", apphandler.ProduceMessage())

	// Run server
	port := os.Getenv("PA_PORT")
	host := "0.0.0.0"
	if port == "" {
		host = host + ":8080"
	} else {
		host = host + ":" + port
	}
	apiServer.Run(host)

}
