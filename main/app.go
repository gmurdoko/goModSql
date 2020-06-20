package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	DB_SCHEMA string
)

func main() {
	DB_USER = "root"
	DB_PASSWORD = "toor"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_SCHEMA = "toko_enigma"
	connDB(&DB_USER, &DB_PASSWORD, &DB_HOST, &DB_PORT, &DB_SCHEMA)

}

func connCheck(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func connDB(DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_SCHEMA *string) {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *DB_USER, *DB_PASSWORD, *DB_HOST, *DB_PORT, *DB_SCHEMA)
	db, _ := sql.Open("mysql", dbSource)
	connCheck(db)
}