package main

import (
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	crudRoutes "golang-rest-api-mysql-with-redis/src/features/crud/presentation/routes"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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

	router := gin.New()

	router.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(200, "index.tmpl", gin.H{
			"hostname": os.Getenv(config.LocalHostname),
			"cores":    os.Getenv(config.LocalCoreCount),
		})
	})

	apiv1 := router.Group("/api/v1")

	crud := apiv1.Group("/crud")

	crudRoutes.LoadRoutes(crud)

	router.Run(":" + os.Getenv(config.Port))

}
