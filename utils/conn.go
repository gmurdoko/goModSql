package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connCheck(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("cie berhasil")
}

func ConnDB(DB_ENGINE, DbSource string) {
	db, _ := sql.Open(DB_ENGINE, DbSource)
	connCheck(db)
}
