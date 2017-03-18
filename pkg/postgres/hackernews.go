package postgres

import (
	"log"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

func NewsExists(id int64) bool {

	var cnt int
	err := db.Get(&cnt, "SELECT count(*) FROM HackerNews WHERE Id = $1", id)

	if err != nil {
		log.Fatal(err)
	}

	return cnt != 0
}

func InsertNews(n quzx_crawler.HackerNews) {

	tx := db.MustBegin()

	insertQuery := `INSERT INTO HackerNews(Id, By, Score, Time, Title, Type, Url, Readed)
		        VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := tx.Exec(insertQuery,
			  n.Id, n.By, n.Score, n.Time, n.Title, n.Type, n.Url, 0)

	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}
