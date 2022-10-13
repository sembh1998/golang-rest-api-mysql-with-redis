package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.RouterGroup) {
	r.GET("/employees", getAllEmployees)
	r.GET("/departments", getAllDepartments)
}
