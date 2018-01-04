package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/demas/cowl-go/pkg/logging"
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
		logging.PostgreLog{}.LogError(err.Error())
	}
}
