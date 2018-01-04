package postgres

import (
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"github.com/demas/cowl-go/pkg/logging"
	"github.com/demas/cowl-go/pkg/rest-api/quzx"
)

// represent a PostgreSQL implementation of quzx_crawler.HackerNewsService
type HackerNewsRepository struct {
}

func (r *HackerNewsRepository) GetHackerNewsById(id int64) (*quzx.HackerNews, error) {

	selectItemsQuery := `SELECT * FROM HackerNews WHERE Id = $1`
	result := quzx.HackerNews{}

	err := db.Get(&result, selectItemsQuery, id)
	if err != nil {
		logging.PostgreLog{}.LogInfo(err.Error())
	}

	return &result, err
}

func (r *HackerNewsRepository) NewsExists(id int64) bool {

	var cnt int
	err := db.Get(&cnt, "SELECT count(*) FROM HackerNews WHERE Id = $1", id)

	if err != nil {
		logging.PostgreLog{}.LogError(err.Error())
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
		logging.PostgreLog{}.LogError(err.Error())
	}

	tx.Commit()
}

func (s *HackerNewsRepository) GetUnreadedHackerNews() ([]*quzx.HackerNews, error) {

	result := []*quzx.HackerNews{}
	err := db.Select(&result, `SELECT * FROM HackerNews WHERE Readed = 0 ORDER BY TIME DESC`)
	return result, err
}

func (s *HackerNewsRepository) SetHackerNewsAsReaded(id int64) {

	tx := db.MustBegin()
	_, err := tx.Exec(`UPDATE HackerNews SET READED = 1 WHERE Id = $1`, id)
	if err != nil {
		logging.PostgreLog{}.LogInfo(err.Error())
	}
	tx.Commit()
}

func (s *HackerNewsRepository) SetHackerNewsAsReadedFromTime(t int64) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE HackerNews SET READED = 1 WHERE Time < $1", t)
	if err != nil {
		logging.PostgreLog{}.LogInfo(err.Error())
	}
	tx.Commit()
}

func (s *HackerNewsRepository) SetAllHackerNewsAsReaded() {

	tx := db.MustBegin()
	_, err := tx.Exec(`UPDATE HackerNews SET READED = 1`)
	if err != nil {
		logging.PostgreLog{}.LogInfo(err.Error())
	}
	tx.Commit()
}

func (r *HackerNewsRepository) DeleteAllHackerNews() {

	deleteHackerNewsQuery := `DELETE FROM HackerNews`

	tx := db.MustBegin()
	tx.MustExec(deleteHackerNewsQuery)
	tx.Commit()
}
