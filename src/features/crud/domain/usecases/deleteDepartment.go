package usecases

import (
	"context"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
)

func DeleteDepartment(id int32) error {
	mysqlconn := config.GetMysqlConnection()
	if mysqlconn == nil {
		return fmt.Errorf("error getting mysql connection")
	}
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	resp, err := db.DeleteDepartment(context.Background(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no department with id %d", id)
	}

	table := "department"

	go cache.UpdateTableGeneric(table)

	return nil
}
