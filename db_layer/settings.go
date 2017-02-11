package db_layer

import (
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
)

var db *sqlx.DB

func init() {
	var err error

	db, err = sqlx.Open("postgres", "user=demas password=root host=192.168.1.71 port=5432 dbname=news sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

type Settings struct {
	Name string
	Value string
}

func GetSettings(key string) string {

	settings := []Settings{}

	err := db.Select(&settings, fmt.Sprintf("SELECT * FROM Settings WHERE Name = '%s'", key))
	if err != nil {
		log.Fatal(err)
	}

	if len(settings) == 0 {
		return ""
	} else {
		return settings[0].Value
	}
}

func SetSettings(key string, value string)  {

	settings := []Settings{}

	err := db.Select(&settings, fmt.Sprintf("SELECT * FROM Settings WHERE Name = '%s'", key))
	if err != nil {
		log.Fatal(err)
	}

	tx := db.MustBegin()
	if len(settings) == 0 {
		tx.MustExec("INSERT INTO Settings(Name, Value) VALUES($1, $2)", key, value)
	} else {
		tx.MustExec("UPDATE Settings SET Value=$2 WHERE Name=$1", key, value)
	}
	tx.Commit()
}