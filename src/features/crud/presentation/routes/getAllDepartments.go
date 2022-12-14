package routes

import (
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/usecases"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getAllDepartments(c *gin.Context) {
	departments, err := usecases.GetAllDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"serving": os.Getenv(config.LocalHostname),
			"cores":   os.Getenv(config.LocalCoreCount),
		})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"departments": departments,
			"serving":     os.Getenv(config.LocalHostname),
			"cores":       os.Getenv(config.LocalCoreCount),
		})
}
