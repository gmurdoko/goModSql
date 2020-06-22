package view

import (
	"fmt"
	"log"
	"mySqlApp/config"
	"mySqlApp/utils"
)

type films struct {
	id           int
	title        string
	releasedYear int
	genre        string
}

func GetAllData() {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM series;")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result []films
	for data.Next() {
		var film = films{}
		var err = data.Scan(&film.id, &film.title, &film.releasedYear, &film.genre)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, film)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func GetDataById() {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows := 1
	var film = films{}

	// data := db.QueryRow("SELECT * FROM series WHERE id = ?;", rows)
	// err = data.Scan(&film.id, &film.title, &film.releasedYear, &film.genre)

	err = db.QueryRow("SELECT * FROM series WHERE id = ?;", rows).Scan(&film.id, &film.title, &film.releasedYear, &film.genre)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(film)
}
func PrepareQuery() {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	findByID, err := db.Prepare("SELECT * FROM series WHERE id =?")
	if err != nil {
		log.Fatal(err)
	}
	defer findByID.Close()
	data, err := findByID.Query(1)
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	var result []films
	for data.Next() {
		var film = films{}
		var err = data.Scan(&film.id, &film.title, &film.releasedYear, &film.genre)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, film)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
