package usecases

import (
	"context"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
)

func UpdateDepartment(department adapter.DepartmentRequest) error {
	mysqlconn := config.GetMysqlConnection()
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	_, err := db.UpdateDepartment(context.Background(), mysqlsimplecrud.UpdateDepartmentParams{
		Name: department.Name,
		ID:   department.ID,
	})
	if err != nil {
		return err
	}
	table := "departments"

	go cache.UpdateTableGeneric(table)

	return nil
}
