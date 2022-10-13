package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.RouterGroup) {
	r.GET("/employees", getAllEmployees)
	r.GET("/departments", getAllDepartments)
	r.GET("/employees/:id", getEmployee)
	r.GET("/departments/:id", getDepartment)

	r.POST("/departments", createDepartment)
	r.PUT("/departments/:id", updateDepartment)
	r.DELETE("/departments/:id", deleteDepartment)

	r.POST("/employees", createEmployee)
	r.PUT("/employees/:id", updateEmployee)
	r.DELETE("/employees/:id", deleteEmployee)

	r.GET("/getEmployeeByDepartment/:name", getEmployeesByDepartment)
}
