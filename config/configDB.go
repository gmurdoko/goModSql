package config

import (
	"fmt"
	"mySqlApp/utils"
)

var (
	//DbEngine,
	DbUser,
	DbPassword,
	DbHost,
	DbPort,
	DbSchema string
)

func Env_conn() (dbEngine, dbSource string) {
	dbEngine = utils.ViperGetEnv("DB_ENGINE", "mysql") //mysql
	DbUser = utils.ViperGetEnv("DB_USER", "root")      //root
	DbPassword = utils.ViperGetEnv("DB_PASSWORD", "password")
	DbHost = utils.ViperGetEnv("DB_HOST", "localhost") //localhost
	DbPort = utils.ViperGetEnv("DB_PORT", "3306")      //3306
	DbSchema = utils.ViperGetEnv("DB_SCHEMA", "schema")

	dbSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbSchema)
	return dbEngine, dbSource
}
