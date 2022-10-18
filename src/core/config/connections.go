package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
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
	Conn *redis.Client
}

var singletonRedis *RedisConnection

func GetRedisConnection() *RedisConnection {
	if singletonRedis == nil {
		db, err := strconv.Atoi(os.Getenv(RedisDB))
		if err != nil {
			panic(err)
		}
		conn := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", os.Getenv(RedisHost), os.Getenv(RedisPort)),
			//Password: os.Getenv(RedisPass),
			//Username: os.Getenv(RedisUser),
			DB: db,
		})
		singletonRedis = &RedisConnection{Conn: conn}
		log.Println("Redis connection created")
	}
	ctx := context.Background()
	timeOutCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	_, err := singletonRedis.Conn.Ping(timeOutCtx).Result()
	if err != nil {
		fmt.Printf("Error connecting to redis: %v\n", err)
		singletonRedis = nil
		return nil
	}
	return singletonRedis
}
