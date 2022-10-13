package routes

import (
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/usecases"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getAllEmployees(c *gin.Context) {
	employees, err := usecases.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"serving": os.Getenv(config.LocalHostname),
			"cores":   os.Getenv(config.LocalCoreCount)})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employees": employees,
		"serving":   os.Getenv(config.LocalHostname),
		"cores":     os.Getenv(config.LocalCoreCount),
	})
}
