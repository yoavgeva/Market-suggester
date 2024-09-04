package main

import (
	"fmt"
	"market-suggester/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
)




func main() {
	var ginServer *gin.Engine
	if environment() == "production" {
		gin.SetMode(gin.ReleaseMode)
		ginServer = gin.New()
	} else {
		ginServer = gin.Default()
	}
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
