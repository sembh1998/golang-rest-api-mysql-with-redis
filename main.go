package main

import (
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	crudRoutes "golang-rest-api-mysql-with-redis/src/features/crud/presentation/routes"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	config.LoadEnvs()
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	os.Setenv(config.LocalHostname, hostname)
	cores := runtime.NumCPU()
	os.Setenv(config.LocalCoreCount, fmt.Sprintf("%d", cores))

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"serving": os.Getenv(config.LocalHostname),
			"cores":   os.Getenv(config.LocalCoreCount),
		})
	})

	apiv1 := router.Group("/api/v1")

	crud := apiv1.Group("/crud")

	crudRoutes.LoadRoutes(crud)

	router.Run(":" + os.Getenv(config.Port))

}
