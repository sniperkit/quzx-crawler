package db_layer

import (
	"log"
	"database/sql"
	"github.com/SlyMarbo/rss"
	"fmt"
)

type RssFeed struct {
	Id int
	Title sql.NullString
	Description sql.NullString
	Link string
	UpdateUrl sql.NullString
	ImageTitle sql.NullString
	ImageUrl sql.NullString
	ImageHeight sql.NullInt64
	ImageWidth sql.NullInt64
	LastSyncTime sql.NullInt64
	Total sql.NullInt64
	Readed sql.NullInt64
}

type RssItem struct {
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

func GetFeeds() []RssFeed {

	feeds := []RssFeed{}

	err := db.Select(&feeds, "SELECT * FROM RssFeed")
	if err != nil {
		log.Fatal(err)
	}

	return feeds
}

func UpdateFeed(id int, feed *rss.Feed, lastSyncTime int64) {

	tx := db.MustBegin()
	tx.MustExec("UPDATE RssFeed SET Title=$1, Description = $2, UpdateUrl = $3, ImageTitle = $4, " +
				"ImageUrl = $5, ImageHeight = $6, ImageWidth = $7, LastSyncTime = $8  WHERE Id=$9",
		feed.Title, feed.Description, feed.UpdateURL, feed.Image.Title, feed.Image.Url, feed.Image.Height,
		feed.Image.Width, lastSyncTime, id)
	tx.Commit()
}

func InsertRssItem(feed_id int, i *rss.Item) {

	items := []RssItem{}

	err := db.Select(&items, fmt.Sprintf("SELECT * FROM RssItem WHERE FeedId = %d AND Link = '%s'", feed_id, i.Link))
	if err != nil {
		log.Fatal(err)
	}

	tx := db.MustBegin()
	if len(items) == 0 {
		tx.MustExec("INSERT INTO RssItem(FeedId, Title, Summary, Content, Link, Date, ItemId, Readed) " +
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8)", feed_id, i.Title, i.Content, i.Summary, i.Link, i.Date.Unix(), i.ID, 0)
	}
	tx.Commit()
}