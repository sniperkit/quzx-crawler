package feed

import (
	"github.com/demas/cowl-go/db_layer"
	"github.com/SlyMarbo/rss"

	"log"
	"time"
)

func Fetch() {

	db_feeds := db_layer.GetFeeds()
	for _, db_feed := range db_feeds {

		f, err := rss.Fetch(db_feed.Link)
		if err != nil {
			log.Println(err)
		}
		db_layer.UpdateFeed(db_feed.Id, f, time.Now().Unix())

		for _, item := range f.Items {
			db_layer.InsertRssItem(db_feed.Id, item)
		}
	}
}
