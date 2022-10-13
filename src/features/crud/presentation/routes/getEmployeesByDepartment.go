package routes

import (
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/usecases"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getEmployeesByDepartment(c *gin.Context) {
	department := c.Param("name")
	employees, err := usecases.GetEmployeesByDepartment(department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"serving": os.Getenv(config.LocalHostname),
			"cores":   os.Getenv(config.LocalCoreCount),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"serving":   os.Getenv(config.LocalHostname),
		"cores":     os.Getenv(config.LocalCoreCount),
		"employees": employees,
	})
}
