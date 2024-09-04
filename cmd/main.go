package main

import (
	"fmt"
	"market-suggester/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
)




func main() {
	ginServer := setupRouter()

	ginServer.GET("/health",handlers.HealthHandler)
	err := ginServer.Run(":8080")
	if err != nil {
		fmt.Println("Server crash")
		panic(err)
	}

}

func environment() string {
	return os.Getenv("ENV")
}

func setupRouter() *gin.Engine {
	var r *gin.Engine
	if environment() == "production" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		r = gin.Default()
	}
	
	return r
}
