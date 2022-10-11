package main

import (
	"fmt"
	"golang-rest-api-mysql-with-redis/core/config"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(os.Getenv(config.Port))

}
