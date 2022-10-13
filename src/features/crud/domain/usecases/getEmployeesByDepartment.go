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

func GetEmployeesByDepartment(department string) ([]adapter.EmployeeResponse, error) {
	consultInternal := fmt.Sprintf("GetDepartmentByName{name:%v}", department)
	tablesInternal := []string{"department"}

	cacheValue, err := cache.GetGeneric(consultInternal, tablesInternal)
	departmentInternal := adapter.DepartmentResponse{}
	if err == nil {
		log.Println("Cache hit")
		err = json.Unmarshal([]byte(cacheValue), &departmentInternal)
		if err != nil {
			log.Println("Error unmarshalling department", departmentInternal)
		}
	} else {
		log.Printf("Error getting dapartment from cache: %v\n", err)
		log.Println("Getting dapartment from database")
		mysqlconn := config.GetMysqlConnection()
		db := mysqlsimplecrud.New(mysqlconn.Conn)
		depInt, err := db.GetDepartmentByName(context.Background(), department)
		if err != nil {
			return nil, err
		}
		departmentInternal = adapter.ToDepartmentResponse(depInt)

		stringValue, err := json.Marshal(departmentInternal)
		if err != nil {
			log.Printf("Error marshalling employee: %v\n", err)
		} else {
			go cache.SetGeneric(consultInternal, string(stringValue))
		}
	}

	consult := "GetDepartmentEmployees{name:" + department + "}"
	tables := []string{"employee", "department"}

	cacheValue, err = cache.GetGeneric(consult, tables)
	if err == nil {
		log.Println("Cache hit")
		returnValue := []adapter.EmployeeResponse{}
		err = json.Unmarshal([]byte(cacheValue), &returnValue)
		if err != nil {
			log.Printf("Error unmarshalling employee: %v\n", err)
			return nil, err
		}
		return returnValue, nil
	}
	log.Printf("Error getting dapartment from cache: %v\n", err)
	log.Println("Getting dapartment from database")

	mysqlconn := config.GetMysqlConnection()

	db := mysqlsimplecrud.New(mysqlconn.Conn)

	employees, err := db.GetDepartmentEmployees(context.Background(), departmentInternal.ID)

	if err != nil {
		return nil, err
	}

	response := adapter.ToGetEmployeesByDepartmentResponse(employees, departmentInternal.Name)

	stringValue, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling employee: %v\n", err)
		return response, nil
	}
	go cache.SetGeneric(consult, string(stringValue))
	return response, nil
}
