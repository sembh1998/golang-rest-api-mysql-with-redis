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

func GetDepartment(id int32) (adapter.DepartmentResponse, error) {
	consult := fmt.Sprintf("GetDepartment{id:%v}", id)

	tables := []string{"departments"}

	cacheValue, err := cache.GetGeneric(consult, tables)
	if err == nil {
		log.Println("Cache hit")
		returnValue := adapter.DepartmentResponse{}
		err = json.Unmarshal([]byte(cacheValue), &returnValue)
		if err != nil {
			return adapter.DepartmentResponse{}, err
		}
		return returnValue, nil
	}
	log.Printf("Error getting dapartment from cache: %v\n", err)
	log.Println("Getting dapartment from database")

	mysqlconn := config.GetMysqlConnection()
	if mysqlconn == nil {
		return adapter.DepartmentResponse{}, fmt.Errorf("error getting mysql connection")
	}
	db := mysqlsimplecrud.New(mysqlconn.Conn)

	department, err := db.GetDepartment(context.Background(), id)

	if err != nil {
		return adapter.DepartmentResponse{}, err
	}
	response := adapter.ToDepartmentResponse(department)

	stringValue, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling employee: %v\n", err)
		return response, nil
	}
	go cache.SetGeneric(consult, string(stringValue))

	return response, nil
}
