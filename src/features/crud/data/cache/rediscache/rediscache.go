package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/core/utils"
	"golang-rest-api-mysql-with-redis/src/features/crud/domain/adapter"
	"strconv"
	"time"
)

var (
	ConsultGetAllDepartments          = "GetAllDepartments"
	GetAllDepartmentsImplicatedTables = []string{"department"}

	ConsultGetAllEmployees          = "GetAllEmployees"
	GetAllEmployeesImplicatedTables = []string{"employee", "department"}
)

func GetAllDepartments() ([]adapter.DepartmentResponse, error) {
	redisconn := config.GetRedisConnection()

	rbd := redisconn.Conn

	consultDateString, err := rbd.Get(context.Background(), utils.GetRedisConsultDateKey(ConsultGetAllDepartments)).Result()
	if err != nil {
		return nil, err
	}
	consultDate, err := strconv.ParseInt(consultDateString, 10, 64)
	if err != nil {
		return nil, err
	}
	tablesDates := make(map[string]int64)

	for _, table := range GetAllDepartmentsImplicatedTables {
		tableDateString, err := rbd.Get(context.Background(), utils.GetMysqlTableKey(table)).Result()
		if err != nil {
			go rbd.Set(context.Background(), utils.GetMysqlTableKey(table), time.Now().UnixNano(), 0).Err()
			return nil, err
		}
		tableDate, err := strconv.ParseInt(tableDateString, 10, 64)
		if err != nil {
			return nil, err
		}
		tablesDates[table] = tableDate
	}

	for _, table := range GetAllDepartmentsImplicatedTables {
		if tablesDates[table] > consultDate {
			return nil, fmt.Errorf("table %s has been modified", table)
		}
	}
	valueString, err := rbd.Get(context.Background(), utils.GetRedisConsultKey(ConsultGetAllDepartments)).Result()
	if err != nil {
		return nil, err
	}
	returnValue := []adapter.DepartmentResponse{}
	err = json.Unmarshal([]byte(valueString), &returnValue)
	if err != nil {
		return nil, err
	}
	return returnValue, nil
}

func SetGetAllDepartments(value []adapter.DepartmentResponse) error {
	redisconn := config.GetRedisConnection()

	rbd := redisconn.Conn

	valueString, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = rbd.Set(context.Background(), utils.GetRedisConsultKey(ConsultGetAllDepartments), valueString, 0).Err()
	if err != nil {
		return err
	}
	err = rbd.Set(context.Background(), utils.GetRedisConsultDateKey(ConsultGetAllDepartments), time.Now().UnixNano(), 0).Err()
	if err != nil {
		return err
	}
	return nil
}
