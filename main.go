package main

import (
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	config.LoadEnvs()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
				"serving": "no hostname",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"serving": hostname,
		})
	})

	router.Run(":" + os.Getenv(config.Port))

}
