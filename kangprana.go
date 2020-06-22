package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id   int
	name string
	age  int
}

func main() {
	getAllData()
}

func connect() (*sql.DB, error) {
	db, _ := sql.Open("mysql", "root:P@ssw0rd@tcp(127.0.0.1:3306)/db_contoh")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func getAllData() {
	db, err := connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	data, err := db.Query("select * from tb_student where id=?", 1)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result []student

	for data.Next() {
		var siswa = student{}
		var err = data.Scan(&siswa.id, &siswa.name, &siswa.age)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, siswa)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
