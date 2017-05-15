package logging

import (
	"log"
	"time"

	"fmt"
	"os"

	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"github.com/jmoiron/sqlx"
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
		log.Fatal(err.Error())
	}
}

func InsertLogMessage(message quzx_crawler.LogMessage) {

	tx := db.MustBegin()

	insertQuery := `INSERT INTO Logs(Moment, Application, Level, Message)
		        VALUES($1, $2, $3, $4)`

	_, err := tx.Exec(insertQuery,
		message.Moment,
		message.Application,
		message.Level,
		message.Message)

	if err != nil {
		log.Println("Error inserting logs to DB")
	}

	tx.Commit()
}

func LogInfo(message string) {

	LogMessage(quzx_crawler.LogMessage{Moment: time.Now().Unix(), Application: "crawler", Level: 5, Message: message})
}

func LogError(message string) {

	LogMessage(quzx_crawler.LogMessage{Moment: time.Now().Unix(), Application: "crawler", Level: 1, Message: message})
}

func LogMessage(message quzx_crawler.LogMessage) {

	InsertLogMessage(message)

	if message.Level == 5 {
		log.Println(message.Message)
	} else if message.Level == 1 {
		log.Fatal(message.Message)
	}
}
