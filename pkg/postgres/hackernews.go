package postgres

import (
	"log"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

// represent a PostgreSQL implementation of quzx_crawler.HackerNewsService
type HackerNewsRepository struct {
}

func (r *HackerNewsRepository) NewsExists(id int64) bool {

	var cnt int
	err := db.Get(&cnt, "SELECT count(*) FROM HackerNews WHERE Id = $1", id)

	if err != nil {
		log.Fatal(err)
	}

	return cnt != 0
}

func (r *HackerNewsRepository) InsertNews(n quzx_crawler.HackerNews) {

	tx := db.MustBegin()

	insertQuery := `INSERT INTO HackerNews(Id, By, Score, Time, Title, Type, Url, Readed, Favorite)
		        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := tx.Exec(insertQuery,
			  n.Id, n.By, n.Score, n.Time, n.Title, n.Type, n.Url, 0, 0)

	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}
