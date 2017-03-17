package db_layer

import (
	"log"
	"github.com/SlyMarbo/rss"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"fmt"
)

func GetFeeds() []quzx_crawler.RssFeed {

	feeds := []quzx_crawler.RssFeed{}

	err := db.Select(&feeds, "SELECT * FROM RssFeed")
	if err != nil {
		log.Fatal(err)
	}

	return feeds
}

func UpdateFeed(id int, feed *rss.Feed, lastSyncTime int64) {

	tx := db.MustBegin()
	tx.MustExec("UPDATE RssFeed SET Title=$1, Description = $2, UpdateUrl = $3, ImageTitle = $4, " +
				"ImageUrl = $5, ImageHeight = $6, ImageWidth = $7, LastSyncTime = $8, Broken = 0  WHERE Id=$9",
		feed.Title, feed.Description, feed.UpdateURL, feed.Image.Title, feed.Image.Url, feed.Image.Height,
		feed.Image.Width, lastSyncTime, id)
	tx.Commit()
}

func UpdateFeedAsBroken(id int) {

	tx := db.MustBegin()
	tx.MustExec("UPDATE RssFeed SET Broken = 1 WHERE Id=$1", id)
	tx.Commit()
}

func InsertRssItem(feed_id int, i *rss.Item) {

	items := []quzx_crawler.RssItem{}

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