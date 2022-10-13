package adapter

import "golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"

type DepartmentResponse struct {
	ID   int32  `json:"department_id"`
	Name string `json:"department_name"`
}

func ToDepartmentResponse(department mysqlsimplecrud.Department) DepartmentResponse {
	return DepartmentResponse{
		ID:   department.ID,
		Name: department.Name,
	}
}

func ToDepartmentResponses(departments []mysqlsimplecrud.Department) []DepartmentResponse {
	departmentResponses := make([]DepartmentResponse, len(departments))
	for i, department := range departments {
		departmentResponses[i] = ToDepartmentResponse(department)
	}
	return departmentResponses
}
