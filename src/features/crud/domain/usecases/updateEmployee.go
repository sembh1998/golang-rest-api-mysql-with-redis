package usecases

import (
	"context"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
)

func UpdateEmployee(employee adapter.EmployeeRequest) error {
	mysqlconn := config.GetMysqlConnection()
	if mysqlconn == nil {
		return fmt.Errorf("error getting mysql connection")
	}
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	_, err := db.UpdateEmployee(context.Background(), mysqlsimplecrud.UpdateEmployeeParams{
		Name:         employee.Name,
		DepartmentID: employee.DepartmentRequest.ID,
		ID:           employee.ID,
	})
	if err != nil {
		return err
	}
	table := "employee"

	go cache.UpdateTableGeneric(table)

	return nil
}
