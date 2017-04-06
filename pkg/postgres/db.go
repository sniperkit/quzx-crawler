package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/demas/cowl-go/pkg/logging"
)

var db *sqlx.DB

func init() {

	var err error

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"))

	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		logging.LogError(err.Error())
	}
}
