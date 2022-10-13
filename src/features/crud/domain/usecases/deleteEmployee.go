package usecases

import (
	"context"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
)

func DeleteEmployee(id int32) error {
	mysqlconn := config.GetMysqlConnection()
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	resp, err := db.DeleteEmployee(context.Background(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no employee with id %d", id)
	}

	table := "employee"

	go cache.UpdateTableGeneric(table)

	return nil
}
