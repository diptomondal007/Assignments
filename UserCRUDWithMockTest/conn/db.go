package conn

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dipto"
	password = "12345"
	dbname   = "userDBNEW"
)

func DB() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password,dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil
	}
	log.Println("Connected to Database!")
	return db
}
