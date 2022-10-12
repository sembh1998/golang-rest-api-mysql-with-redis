package main

import (
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	config.LoadEnvs()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		cores := runtime.NumCPU()
		if err != nil {
			hostname = "unknown"
		}
		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"serving": hostname,
			"cores":   cores,
		})
	})

	router.Run(":" + os.Getenv(config.Port))

}
