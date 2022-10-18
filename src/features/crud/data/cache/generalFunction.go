package cache

import (
	"context"
	"fmt"
	"golang-rest-api-mysql-with-redis/src/core/config"
	"golang-rest-api-mysql-with-redis/src/core/utils"
	"strconv"
	"time"
)

func GetGeneric(consult string, relatedTables []string) (string, error) {

	redisconn := config.GetRedisConnection()

	if redisconn == nil {
		return "", fmt.Errorf("redis connection is nil")
	}

	rbd := redisconn.Conn

	consultDateString, err := rbd.Get(context.Background(), utils.GetRedisConsultDateKey(consult)).Result()
	if err != nil {
		return "", err
	}
	consultDate, err := strconv.ParseInt(consultDateString, 10, 64)
	if err != nil {
		return "", err
	}
	tablesDates := make(map[string]int64)

	for _, table := range relatedTables {
		tableDateString, err := rbd.Get(context.Background(), utils.GetMysqlTableKey(table)).Result()
		if err != nil {
			go rbd.Set(context.Background(), utils.GetMysqlTableKey(table), time.Now().UnixNano(), 0).Err()
			return "", err
		}
		tableDate, err := strconv.ParseInt(tableDateString, 10, 64)
		if err != nil {
			return "", err
		}
		tablesDates[table] = tableDate
	}

	for _, table := range relatedTables {
		if tablesDates[table] > consultDate {
			return "", fmt.Errorf("table %s has been modified", table)
		}
	}
	valueString, err := rbd.Get(context.Background(), utils.GetRedisConsultKey(consult)).Result()
	if err != nil {
		return "", err
	}

	return valueString, nil
}

func SetGeneric(consulta, value string) error {
	redisconn := config.GetRedisConnection()

	if redisconn == nil {
		return fmt.Errorf("redis connection is nil")
	}

	rbd := redisconn.Conn

	err := rbd.Set(context.Background(), utils.GetRedisConsultKey(consulta), value, 0).Err()
	if err != nil {
		return err
	}
	err = rbd.Set(context.Background(), utils.GetRedisConsultDateKey(consulta), time.Now().UnixNano(), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func UpdateTableGeneric(table string) error {
	redisconn := config.GetRedisConnection()

	if redisconn == nil {
		return fmt.Errorf("redis connection is nil")
	}

	rbd := redisconn.Conn

	err := rbd.Set(context.Background(), utils.GetMysqlTableKey(table), time.Now().UnixNano(), 0).Err()
	if err != nil {
		return err
	}
	return nil
}
