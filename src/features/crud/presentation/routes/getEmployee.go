package routes

import (
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/usecases"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEmployee(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"serving": os.Getenv(config.LocalHostname),
			"cores":   os.Getenv(config.LocalCoreCount),
		})
		return
	}
	employee, err := usecases.GetEmployee(int32(intid))
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
			"employee": employee,
			"serving":  os.Getenv(config.LocalHostname),
			"cores":    os.Getenv(config.LocalCoreCount),
		})
}
