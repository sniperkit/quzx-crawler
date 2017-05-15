package tst

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

var feedId int

func TestMain(m *testing.M) {

	prepare()
	retCode := m.Run()
	os.Exit(retCode)
}

func prepare() {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"))

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// clean up all tables
	db.Exec(`DELETE FROM Settings`)
	db.Exec(`INSERT INTO Settings(Name, Value) VALUES('one', 'one_value')`)

}
