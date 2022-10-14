package usecases

import (
	"context"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
)

func CreateDepartment(department adapter.DepartmentRequest) (int64, error) {
	mysqlconn := config.GetMysqlConnection()
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	resp, err := db.CreateDepartment(context.Background(), department.Name)
	if err != nil {
		return 0, err
	}
	lastID, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	table := "department"

	go cache.UpdateTableGeneric(table)

	return lastID, nil
}
