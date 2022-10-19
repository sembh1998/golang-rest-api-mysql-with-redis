package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	redis6 "github.com/go-redis/redis/v8"
	redis7 "github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnection struct {
	Conn *sql.DB
}

var singleton *MysqlConnection

func getMySQLStringConnection() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv(MysqlUser),
		os.Getenv(MysqlPass),
		os.Getenv(MysqlHost),
		os.Getenv(MysqlPort),
		os.Getenv(MysqlDB))

	return dataBase
}

func GetMysqlConnection() *MysqlConnection {
	if singleton == nil {
		conn, err := sql.Open("mysql", getMySQLStringConnection())
		if err != nil {
			panic(err)
		}
		singleton = &MysqlConnection{Conn: conn}
		log.Println("Mysql connection created")
	}

	ctx := context.Background()
	timeOutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	err := singleton.Conn.PingContext(timeOutCtx)
	if err != nil {
		fmt.Printf("Error connecting to mysql: %v\n", err)
		singleton = nil
		return nil
	}
	return singleton
}

func (m *MysqlConnection) Close() {
	m.Conn.Close()
	singleton = nil
	log.Println("Mysql connection closed")
}

type RedisConnection struct {
	Conn *redis7.Client
}

var singletonRedis *RedisConnection

func GetRedisConnection() *RedisConnection {
	if singletonRedis == nil {
		db, err := strconv.Atoi(os.Getenv(RedisDB))
		if err != nil {
			panic(err)
		}
		conn := redis7.NewClient(&redis7.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv(RedisHost), os.Getenv(RedisPort)),
			Password: os.Getenv(RedisPass),
			Username: os.Getenv(RedisUser),
			DB:       db,
		})
		singletonRedis = &RedisConnection{Conn: conn}
		log.Println("Redis connection created")
	}
	ctx := context.Background()
	timeOutCtx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()
	res, err := singletonRedis.Conn.Keys(timeOutCtx, "*").Result()
	log.Printf("Redis connection keys: %v\n", res)
	if err != nil {
		fmt.Printf("Error connecting to redis: %v\n", err)
		singletonRedis = nil
		return nil
	}
	return singletonRedis
}

type RedisCacheAws struct {
	Conn *redis6.Client
}

var singletonRedisAws *RedisCacheAws

func GetRedisCacheAws() *RedisCacheAws {
	if singletonRedisAws == nil {
		db, err := strconv.Atoi(os.Getenv(RedisDB))
		if err != nil {
			panic(err)
		}
		conn := redis6.NewClient(&redis6.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv(RedisHost), os.Getenv(RedisPort)),
			Password: os.Getenv(RedisPass),
			Username: os.Getenv(RedisUser),
			DB:       db,
		})
		singletonRedisAws = &RedisCacheAws{Conn: conn}
		log.Println("Redis aws connection created")
	}
	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	res, err := singletonRedisAws.Conn.Keys(ctx, "*").Result()
	log.Printf("Redis aws connection keys: %v\n", res)
	if err != nil {
		fmt.Printf("Error connecting to redis aws: %v\n", err)
		singletonRedisAws = nil
		return nil
	}
	return singletonRedisAws
}
