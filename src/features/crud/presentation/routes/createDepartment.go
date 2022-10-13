package routes

import (
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/usecases"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func createDepartment(c *gin.Context) {
	var department adapter.DepartmentRequest
	err := c.BindJSON(&department)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"serving": os.Getenv(config.LocalHostname),
			"cores":   os.Getenv(config.LocalCoreCount),
		})
		return
	}
	id, err := usecases.CreateDepartment(department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"serving": os.Getenv(config.LocalHostname),
			"cores":   os.Getenv(config.LocalCoreCount),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"serving": os.Getenv(config.LocalHostname),
		"cores":   os.Getenv(config.LocalCoreCount),
		"id":      id,
	})
}
