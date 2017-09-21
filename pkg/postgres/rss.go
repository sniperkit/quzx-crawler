package postgres

import (
	"github.com/SlyMarbo/rss"
	"github.com/demas/cowl-go/pkg/logging"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"github.com/demas/cowl-go/pkg/rest-api/quzx"
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

func (r *RssFeedRepository) GetUnreadRssFeeds(rssType int) ([]*quzx.RssFeed, error) {

	selectQuery := `SELECT * FROM RssFeed WHERE RssType = $1 AND Unreaded > 0`
	result := []*quzx.RssFeed{}
	err := db.Select(&result, selectQuery, rssType)
	return result, err
}

func (s *RssFeedRepository) GetRssFeedById(id int) (quzx.RssFeed, error) {

	selectQuery := `SELECT * FROM RssFeed WHERE Id = $1`
	var result quzx.RssFeed
	err := db.Get(&result, selectQuery, id)
	return result, err
}

func (s *RssFeedRepository) GetRssFeedByUrl(url string) (quzx.RssFeed, error) {

	selectQuery := `SELECT * FROM RssFeed WHERE Link = $1`
	var result quzx.RssFeed
	err := db.Get(&result, selectQuery, url)
	return result, err
}

func (s *RssFeedRepository) InsertRssFeed(feed *quzx.RssFeed) int {

	insertQuery := `INSERT INTO RssFeed
				(Title, Description, Link, UpdateUrl, ImageTitle, ImageUrl, ImageHeight,
				 ImageWidth, LastSyncTime, Total, Unreaded, SyncInterval,  AlternativeName, RssType,
				 ShowContent, ShowOrder, Folder, LimitFull, LimitHeadersOnly, Broken)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
			RETURNING Id`

	tx := db.MustBegin()
	rows, err := tx.Query(insertQuery,
		"",
		"",
		feed.Link,
		"",
		"",
		"",
		0,
		0,
		0,
		0,
		0,
		feed.SyncInterval,
		feed.AlternativeName,
		feed.RssType,
		feed.ShowContent,
		feed.ShowOrder,
		feed.Folder,
		feed.LimitFull,
		feed.LimitHeadersOnly,
		feed.Broken)

	if err != nil {
		logging.LogInfo(err.Error())
	}

	var feedId int = 0
	if rows.Next() {
		rows.Scan(&feedId)
	}

	tx.Commit()

	return feedId
}

func (r *RssFeedRepository) UpdateFeedBeforeSync(title string, description string, updateUrl string, imageTitle string,
	imageUrl string, imageHeight uint32, imageWidth uint32, lastSyncTime int64, id int) {

	tx := db.MustBegin()

	updateQuery := `UPDATE RssFeed
	                SET Title=$1, Description = $2, UpdateUrl = $3, ImageTitle = $4,
			    ImageUrl = $5, ImageHeight = $6, ImageWidth = $7, LastSyncTime = $8, Broken = 0
			WHERE Id=$9`

	tx.MustExec(updateQuery,
		title,
		description,
		imageUrl,
		imageTitle,
		imageUrl,
		imageHeight,
		imageWidth,
		lastSyncTime,
		id)
	tx.Commit()
}

func (s *RssFeedRepository) UpdateFeedAfterSync(link string, lastSyncTime int64, alternativeName string, rssType int,
	showContent int, showOrder int, folder string, syncInterval int, id int) {

	tx := db.MustBegin()

	updateQuery := `UPDATE RssFeed
	                SET Link = $1, LastSyncTime = $2, AlternativeName = $3, RssType = $4,
	                    ShowContent = $5, ShowOrder = $6, Folder = $7, SyncInterval = $8
	                WHERE Id = $9`

	_, err := tx.Exec(updateQuery,
		link,
		lastSyncTime,
		alternativeName,
		rssType,
		showContent,
		showOrder,
		folder,
		syncInterval,
		id)

	if err != nil {
		logging.LogInfo(err.Error())
	}
	tx.Commit()
}

func (r *RssFeedRepository) SetFeedAsBroken(id int) {

	tx := db.MustBegin()
	tx.MustExec("UPDATE RssFeed SET Broken = 1 WHERE Id=$1", id)
	tx.Commit()
}

func (s *RssFeedRepository) SetRssFeedAsReaded(feedId int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE RssItem SET READED = 1 WHERE FeedId = $1", feedId)
	if err != nil {
		logging.LogInfo(err.Error())
	}
	tx.Commit()
}

func (s *RssFeedRepository) UnsubscribeRssFeed(feedId int) {

	tx := db.MustBegin()
	_, err := tx.Exec("DELETE FROM RssItem WHERE FeedId = $1", feedId)
	if err != nil {
		logging.LogInfo(err.Error())
	}
	_, err = tx.Exec("DELETE FROM RssFeed WHERE Id = $1", feedId)
	if err != nil {
		logging.LogInfo(err.Error())
	}
	tx.Commit()
}

func (s *RssFeedRepository) GetRssItemsByFeedId(feed_id int) ([]*quzx.RssItem, error) {

	var feed quzx.RssFeed
	feed, err := s.GetRssFeedById(feed_id)
	if err != nil {
		logging.LogError(err.Error())
	}

	strOrder := feed.OrderByClause()
	limit := feed.Limit()
	selectItemsQuery := `SELECT * FROM RssItem WHERE FeedId = $1 AND Readed = 0 ` + strOrder + ` LIMIT $2`
	result := []*quzx.RssItem{}

	err = db.Select(&result, selectItemsQuery, feed_id, limit)
	if err != nil {
		logging.LogInfo(err.Error())
	}

	return result, err
}

func (s *RssFeedRepository) GetRssItemById(id int) (*quzx.RssItem, error) {

	selectItemsQuery := `SELECT * FROM RssItem WHERE Id = $1`
	result := quzx.RssItem{}

	err := db.Get(&result, selectItemsQuery, id)
	if err != nil {
		logging.LogInfo(err.Error())
	}

	return &result, err
}

func (r *RssFeedRepository) InsertRssItem(feed_id int, i *rss.Item) int {

	insertQuery := `INSERT INTO RssItem(FeedId, Title, Summary, Content, Link, Date, ItemId, Readed, Favorite)
	                VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	                ON CONFLICT (FeedId, Link) DO NOTHING
	                RETURNING Id`

	tx := db.MustBegin()

	rows, err := tx.Query(insertQuery,
		feed_id,
		i.Title,
		i.Content,
		i.Summary,
		i.Link,
		i.Date.Unix(),
		i.ID,
		0,
		0)

	if err != nil {
		logging.LogInfo(err.Error())
		tx.Rollback()
		return 0
	} else {

		var itemId int = 0
		if rows.Next() {
			rows.Scan(&itemId)
		}

		tx.Commit()

		return itemId
	}
}

func (s *RssFeedRepository) SetRssItemAsReaded(id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE RssItem SET READED = 1 WHERE Id = $1", id)
	if err != nil {
		logging.LogInfo(err.Error())
	}
	tx.Commit()
}

func (r *RssFeedRepository) DeleteAllRssFeeds() {

	deleteRssItemsQuery := `DELETE FROM RssItem`
	deleteRssFeedsQuery := `DELETE FROM RssFeed`

	tx := db.MustBegin()
	tx.MustExec(deleteRssItemsQuery)
	tx.MustExec(deleteRssFeedsQuery)
	tx.Commit()
}
