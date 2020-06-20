package main

import (
	"fmt"
	"mySqlApp/utils"
	"os"
)

var (
	DB_ENGINE,
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	DB_SCHEMA string
)

func main() {
	DB_ENGINE = os.Getenv("dbEngine") //mysql
	DB_USER = os.Getenv("dbUser")     //root
	DB_PASSWORD = os.Getenv("dbPass") //toor
	DB_HOST = os.Getenv("dbHost")     //localhost
	DB_PORT = os.Getenv("dbPort")     //3306
	DB_SCHEMA = os.Getenv("dbSchema") //toko_enigma

	DbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_SCHEMA)
	utils.ConnDB(DB_ENGINE, DbSource)
}
