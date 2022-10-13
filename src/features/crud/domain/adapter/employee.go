package adapter

type EmployeeRequest struct {
	Name string `json:"employee_name"`
	ID   int32  `json:"employee_id"`
	DepartmentRequest
}
