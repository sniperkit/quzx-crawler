package postgres

import (
	"fmt"
	"log"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

// GetSettings : return settings from database by key
func GetSettings(key string) string {

	settings :=  quzx_crawler.Settings{}
	query := fmt.Sprintf("SELECT * FROM Settings WHERE Name = '%s' LIMIT 1", key)

	err := db.Get(&settings, query)
	if err != nil {
		log.Println(fmt.Sprintf("Error while get Settings %s : %s", key, err))
	}

	return settings.Value
}

// SetSettings : put setting to database
func SetSettings(key string, value string) {

	settings := []quzx_crawler.Settings{}

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
