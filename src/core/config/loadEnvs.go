package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvs loads the environment variables from the .env file
func LoadEnvs() {
	validateIfEnvVarFileExists("backend")

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

	validateEnvVars(Port, "8084")

}

func validateIfEnvVarFileExists(filename string) {
	path := fmt.Sprintf("envs/%s.env", filename)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Printf(" %s.env file not found\n", filename)
		return
	}
	err = godotenv.Load(path)
	if err != nil {
		log.Printf("File found\n but some syntaxis error loading %s.env file\n", filename)
		return
	}
}

func validateEnvVars(envVar string, defaultValue string) {
	_, err := os.LookupEnv(envVar)
	if !err {
		fmt.Printf("Environment variable '%s' not found, setting default value\n", envVar)
		os.Setenv(envVar, defaultValue)
		return
	}
	fmt.Printf("Environment variable '%s' found\n", envVar)
}
