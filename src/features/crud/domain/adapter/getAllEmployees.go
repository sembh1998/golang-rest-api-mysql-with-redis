package adapter

import "golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"

type EmployeeResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DepartmentResponse
}

func ToEmployeeResponse(employee mysqlsimplecrud.GetAllEmployeesRow) EmployeeResponse {
	return EmployeeResponse{
		ID:   employee.ID,
		Name: employee.Name,
		DepartmentResponse: DepartmentResponse{
			ID:   employee.DepartmentID,
			Name: employee.Department.String,
		},
	}
}

func ToEmployeeResponses(employees []mysqlsimplecrud.GetAllEmployeesRow) []EmployeeResponse {
	employeeResponses := make([]EmployeeResponse, len(employees))
	for i, employee := range employees {
		employeeResponses[i] = ToEmployeeResponse(employee)
	}
	return employeeResponses
}
