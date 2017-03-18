package postgres

import (
	"fmt"
	"log"

	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

// GetSettings : return settings from database by key
func GetSettings(key string) string {

	settings := quzx_crawler.Settings{}
	query := fmt.Sprintf("SELECT * FROM Settings WHERE Name = '%s' LIMIT 1", key)

	err := db.Get(&settings, query)
	if err != nil {
		log.Println(fmt.Sprintf("Error while get Settings %s : %s", key, err))
	}

	return settings.Value
}

// SetSettings : put setting to database
func SetSettings(key string, value string) {

	query := `INSERT INTO Settings(Name, Value) VALUES($1, $2)
			  ON CONFLICT (Name) DO Update SET Value = Excluded.Value`

	tx := db.MustBegin()

	_, err := tx.Exec(query, key, value)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}
