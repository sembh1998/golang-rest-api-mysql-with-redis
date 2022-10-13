package utils

var (
	RedisPrefix            = "redis."
	ConsultDateTermination = ".cache.UnixNanoDate"

	MysqlPrefix      = "mysql."
	TableTermination = ".table.LastChange.UnixNanoDate"
)

func GetRedisConsultDateKey(key string) string {
	return RedisPrefix + key + ConsultDateTermination
}

func GetRedisConsultKey(key string) string {
	return RedisPrefix + key
}

func GetMysqlTableKey(key string) string {
	return MysqlPrefix + key + TableTermination
}
