package db_layer

import (
	"log"
	"database/sql"
	"github.com/SlyMarbo/rss"
	"fmt"
)

type Reddit struct {
	Id int
	Title sql.NullString
	Description sql.NullString
	Link string
	LastSyncTime int64
	Total sql.NullInt64
	Unreaded sql.NullInt64
	SyncInterval int
	AlternativeName string
}

type RedditItem struct {
	Id int
	FeedId int
	Title string
	Summary string
	Content string
	Link string
	Date int64
	ItemId string
	Readed int
}

func GetReddits() []Reddit {

	reddits := []Reddit{}

	err := db.Select(&reddits, "SELECT * FROM Reddit")
	if err != nil {
		log.Fatal(err)
	}

	return reddits
}

func UpdateReddit(id int, feed *rss.Feed, lastSyncTime int64) {

	tx := db.MustBegin()
	tx.MustExec("UPDATE Reddit SET Title=$1, Description = $2, LastSyncTime = $3  WHERE Id= $4",
		feed.Title, feed.Description, lastSyncTime, id)
	tx.Commit()
}

func InsertRedditItem(feed_id int, i *rss.Item) {

	items := []RedditItem{}

	err := db.Select(&items, fmt.Sprintf("SELECT * FROM RedditItem WHERE FeedId = %d AND Link = '%s'", feed_id, i.Link))
	if err != nil {
		log.Fatal(err)
	}

	tx := db.MustBegin()
	if len(items) == 0 {
		tx.MustExec("INSERT INTO RedditItem(FeedId, Title, Summary, Content, Link, Date, ItemId, Readed) " +
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8)", feed_id, i.Title, i.Content, i.Summary, i.Link, i.Date.Unix(), i.ID, 0)
	}
	tx.Commit()
}