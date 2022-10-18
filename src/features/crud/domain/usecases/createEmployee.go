package usecases

import (
	"context"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
)

func CreateEmployee(employee adapter.EmployeeRequest) (int64, error) {
	mysqlconn := config.GetMysqlConnection()
	if mysqlconn == nil {
		return 0, fmt.Errorf("error getting mysql connection")
	}
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	resp, err := db.CreateEmployee(context.Background(), mysqlsimplecrud.CreateEmployeeParams{
		Name:         employee.Name,
		DepartmentID: employee.DepartmentRequest.ID,
	})
	if err != nil {
		return 0, err
	}
	lastID, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	table := "employee"

	go cache.UpdateTableGeneric(table)

	return lastID, nil
}
