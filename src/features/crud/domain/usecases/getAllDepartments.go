package usecases

import (
	"context"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache/rediscache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
	"log"
)

func GetAllDepartments() ([]adapter.DepartmentResponse, error) {

	cacheValue, err := rediscache.GetAllDepartments()
	if err == nil {
		log.Println("Cache hit")
		return cacheValue, nil
	}
	log.Printf("Error getting departments from cache: %v\n", err)
	log.Println("Getting departments from database")

	mysqlconn := config.GetMysqlConnection()

	db := mysqlsimplecrud.New(mysqlconn.Conn)

	departments, err := db.GetAllDepartments(context.Background())

	if err != nil {
		return nil, err
	}

	reponse := adapter.ToDepartmentResponses(departments)

	go rediscache.SetGetAllDepartments(reponse)

	return reponse, nil
}
