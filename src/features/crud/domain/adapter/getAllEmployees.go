package adapter

import "golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"

type EmployeeResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DepartmentResponse
}

func ToEmployeeResponse(employee mysqlsimplecrud.GetEmployeeRow) EmployeeResponse {
	return EmployeeResponse{
		ID:   employee.ID,
		Name: employee.Name,
		DepartmentResponse: DepartmentResponse{
			ID:   employee.DepartmentID,
			Name: employee.Department.String,
		},
	}
}

func ToGetAllEmployeesResponse(employee mysqlsimplecrud.GetAllEmployeesRow) EmployeeResponse {
	return EmployeeResponse{
		ID:   employee.ID,
		Name: employee.Name,
		DepartmentResponse: DepartmentResponse{
			ID:   employee.DepartmentID,
			Name: employee.Department.String,
		},
	}
}

func ToGetAllEmployeeResponses(employees []mysqlsimplecrud.GetAllEmployeesRow) []EmployeeResponse {
	employeeResponses := make([]EmployeeResponse, len(employees))
	for i, employee := range employees {
		employeeResponses[i] = ToGetAllEmployeesResponse(employee)
	}
	return employeeResponses
}

func ToGetEmployeesByDepartmentResponse(employee []mysqlsimplecrud.Employee, departmentName string) []EmployeeResponse {
	employeesResponse := make([]EmployeeResponse, len(employee))
	for i, emp := range employee {
		employeesResponse[i] = EmployeeResponse{
			ID:   emp.ID,
			Name: emp.Name,
			DepartmentResponse: DepartmentResponse{
				ID:   emp.DepartmentID,
				Name: departmentName,
			},
		}
	}
	return employeesResponse
}
