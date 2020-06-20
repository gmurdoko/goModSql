func main() {

	db, err := sql.Open("mysql", "root:@Anggaadji13@tcp(localhost:3306)/enigmaresto")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	PingDB(db)

}
func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("Sucess DB Connected")
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}