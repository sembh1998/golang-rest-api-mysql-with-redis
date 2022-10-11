package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvs loads the environment variables from the .env file
func LoadEnvs() {
	validateIfEnvVarFileExists("local")

	validateEnvVars(MysqlPort, "3307")
	validateEnvVars(MysqlHost, "127.0.0.1")
	validateEnvVars(MysqlUser, "developer")
	validateEnvVars(MysqlPass, "password")
	validateEnvVars(MysqlDB, "database")

	validateEnvVars(RedisHost, "localhost")
	validateEnvVars(RedisPort, "6379")
	validateEnvVars(RedisPass, "superdupersecretpasswordnooneknows")
	validateEnvVars(RedisDB, "0")
	validateEnvVars(RedisUser, "")

	validateEnvVars(Port, "8080")

}

func validateIfEnvVarFileExists(filename string) {
	path := fmt.Sprintf("envs/%s.env", filename)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalf(" %s.env file not found", filename)
		return
	}
	err = godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading %s.env file", filename)
		return
	}
}

func validateEnvVars(envVar string, defaultValue string) {
	_, err := os.LookupEnv(envVar)
	if !err {
		fmt.Printf("Environment variable %s not found, setting default value %s", envVar, defaultValue)
		os.Setenv(envVar, defaultValue)
	}
	fmt.Printf("Environment variable %s found, value %s", envVar, os.Getenv(envVar))
}
