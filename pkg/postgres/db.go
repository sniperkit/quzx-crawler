package db_layer

import (
	"github.com/jmoiron/sqlx"
	"os"
	"log"
)

var db *sqlx.DB

func init() {
	var err error

	db, err = sqlx.Open("postgres", "user=" + os.Getenv("DBUSER") +
		" password=" + os.Getenv("DBPASS") + " host=" + os.Getenv("DBHOST") +
		" port=" + os.Getenv("DBPORT") + " dbname=" + os.Getenv("DBNAME") + " sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
