package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/cache"
	"golang-rest-api-mysql-with-redis/src/features/crud/data/sql/mysqlsimplecrud"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
	"log"
)

func GetAllDepartments() ([]adapter.DepartmentResponse, error) {

	consult := "GetAllDepartments"

	tables := []string{"department"}

	cacheValue, err := cache.GetGeneric(consult, tables)
	if err == nil {
		log.Println("Cache hit")
		returnValue := []adapter.DepartmentResponse{}
		err = json.Unmarshal([]byte(cacheValue), &returnValue)
		if err != nil {
			return nil, err
		}
		return returnValue, nil
	}
	log.Printf("Error getting departments from cache: %v\n", err)
	log.Println("Getting departments from database")

	mysqlconn := config.GetMysqlConnection()
	if mysqlconn == nil {
		return nil, fmt.Errorf("error getting mysql connection")
	}

	db := mysqlsimplecrud.New(mysqlconn.Conn)

	departments, err := db.GetAllDepartments(context.Background())

	if err != nil {
		return nil, err
	}

	response := adapter.ToDepartmentResponses(departments)

	stringValue, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling employees: %v\n", err)
		return response, nil
	}
	go cache.SetGeneric(consult, string(stringValue))

	return response, nil
}
