package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sniperkit/quzx-crawler/pkg/logging"
	"github.com/sniperkit/quzx-crawler/pkg/quzxutil"
)

var db *sqlx.DB
var grm *gorm.DB

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

	grm, err = gorm.Open("postgres", "host=192.168.1.71 user=root dbname=rss sslmode=disable password=root")
	if err != nil {
		logging.PostgreLog{}.LogError(err.Error())
	}
}
