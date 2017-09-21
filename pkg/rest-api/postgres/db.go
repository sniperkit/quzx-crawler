package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"github.com/demas/cowl-go/pkg/quzxutil"
)

var db *sqlx.DB

func init() {

	var err error

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		quzxutil.GetParameter("DBUSER"),
		quzxutil.GetParameter("DBPASS"),
		quzxutil.GetParameter("DBHOST"),
		quzxutil.GetParameter("DBPORT"),
		quzxutil.GetParameter("DBNAME"))

	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}