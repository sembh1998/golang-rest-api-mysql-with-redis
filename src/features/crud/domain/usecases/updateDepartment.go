package usecases

import (
	"context"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
)

func UpdateDepartment(department adapter.DepartmentRequest) error {
	mysqlconn := config.GetMysqlConnection()
	if mysqlconn == nil {
		return fmt.Errorf("error getting mysql connection")
	}
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	_, err := db.UpdateDepartment(context.Background(), mysqlsimplecrud.UpdateDepartmentParams{
		Name: department.Name,
		ID:   department.ID,
	})
	if err != nil {
		return err
	}
	table := "department"

	go cache.UpdateTableGeneric(table)

	return nil
}
