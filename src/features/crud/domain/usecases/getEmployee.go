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

func GetEmployee(id int32) (adapter.EmployeeResponse, error) {

	consult := fmt.Sprintf("GetEmployee{id:%v}", id)

	tables := []string{"employee", "department"}

	cacheValue, err := cache.GetGeneric(consult, tables)
	if err == nil {
		log.Println("Cache hit")
		returnValue := adapter.EmployeeResponse{}
		err = json.Unmarshal([]byte(cacheValue), &returnValue)
		if err != nil {
			return adapter.EmployeeResponse{}, err
		}
		return returnValue, nil
	}
	log.Printf("Error getting employee from cache: %v\n", err)
	log.Println("Getting employee from database")

	mysqlconn := config.GetMysqlConnection()
	if mysqlconn == nil {
		return adapter.EmployeeResponse{}, fmt.Errorf("error getting mysql connection")
	}
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	employee, err := db.GetEmployee(context.Background(), id)

	if err != nil {
		return adapter.EmployeeResponse{}, err
	}
	response := adapter.ToEmployeeResponse(employee)

	stringValue, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling employee: %v\n", err)
		return response, nil
	}
	go cache.SetGeneric(consult, string(stringValue))

	return response, nil
}
