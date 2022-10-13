package adapter

type DepartmentRequest struct {
	Name string `json:"department_name"`
	ID   int32  `json:"department_id"`
}
