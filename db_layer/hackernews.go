package db_layer

import (
	"fmt"
	"log"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

func NewsDoesntExists(id int64) bool {

	var cnt int
	err := db.Get(&cnt, fmt.Sprintf("SELECT count(*) FROM HackerNews WHERE Id = '%d'", id))
	if err != nil {
		log.Fatal(err)
	}

	return cnt == 0
}

func InsertNews(n quzx_crawler.HackerNews) {

	tx := db.MustBegin()
	_, err := tx.Exec("INSERT INTO HackerNews(" +
		"Id, By, Score, Time, Title, Type, Url, Readed) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
		n.Id, n.By, n.Score, n.Time, n.Title, n.Type, n.Url, 0)

	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}
