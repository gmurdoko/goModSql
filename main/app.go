package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_SCHEMA)
	connDB(&DB_ENGINE, &dbSource)
}

func connCheck(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("cie berhasil")
}

func connDB(DB_ENGINE, dbSource *string) {
	db, _ := sql.Open(*DB_ENGINE, *dbSource)
	connCheck(db)
}