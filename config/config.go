package config

import (
	"fmt"
	"mySqlApp/utils"
)

var (
	DbEngine,
	DbUser,
	DbPassword,
	DbHost,
	DbPort,
	DbSchema string
)

func Env_conn() {
	DbEngine = utils.ViperGetEnv("DB_ENGINE", "mysql")       //mysql
	DbUser = utils.ViperGetEnv("DB_USER", "root")            //root
	DbPassword = utils.ViperGetEnv("DB_PASSWORD", "toor")    //toor
	DbHost = utils.ViperGetEnv("DB_HOST", "localhost")       //localhost
	DbPort = utils.ViperGetEnv("DB_PORT", "3306")            //3306
	DbSchema = utils.ViperGetEnv("DB_SCHEMA", "toko_enigma") //toko_enigma

	DbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbSchema)
	utils.ConnDB(DbEngine, DbSource)
}
