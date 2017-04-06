package postgres

import (
	"github.com/SlyMarbo/rss"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"github.com/demas/cowl-go/pkg/logging"
)

// represent a PostgreSQL implementation of quzx_crawler.RssFeedRepository
type RssFeedRepository struct {
}


func (r *RssFeedRepository) GetFeeds() []quzx_crawler.RssFeed {

	feeds := []quzx_crawler.RssFeed{}

	err := db.Select(&feeds, "SELECT * FROM RssFeed")
	if err != nil {
		logging.LogError(err.Error())
	}

	return feeds
}

func (r *RssFeedRepository) UpdateFeed(id int, feed *rss.Feed, lastSyncTime int64) {

	tx := db.MustBegin()

	updateQuery := `UPDATE RssFeed
	                SET Title=$1, Description = $2, UpdateUrl = $3, ImageTitle = $4,
			    ImageUrl = $5, ImageHeight = $6, ImageWidth = $7, LastSyncTime = $8, Broken = 0
			WHERE Id=$9`

	tx.MustExec(updateQuery,
		    	feed.Title,
			feed.Description,
			feed.UpdateURL,
			feed.Image.Title,
			feed.Image.Url,
			feed.Image.Height,
			feed.Image.Width,
			lastSyncTime,
			id)
	tx.Commit()
}

func (r *RssFeedRepository) SetFeedAsBroken(id int) {

	tx := db.MustBegin()
	tx.MustExec("UPDATE RssFeed SET Broken = 1 WHERE Id=$1", id)
	tx.Commit()
}

func (r *RssFeedRepository) InsertRssItem(feed_id int, i *rss.Item) {

	insertQuery := `INSERT INTO RssItem(FeedId, Title, Summary, Content, Link, Date, ItemId, Readed, Favorite)
	                VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	                ON CONFLICT (FeedId, Link) DO NOTHING`

	tx := db.MustBegin()
	tx.MustExec(insertQuery,
		feed_id,
		i.Title,
		i.Content,
		i.Summary,
		i.Link,
		i.Date.Unix(),
		i.ID,
		0,
		0)
	tx.Commit()
}