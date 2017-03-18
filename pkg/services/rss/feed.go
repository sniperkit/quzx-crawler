package rssfeeds

import (
	"github.com/demas/cowl-go/pkg/postgres"
	"github.com/SlyMarbo/rss"

	"log"
	"time"
	"net/http"
)

func fetchFunc(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("reddit: cannot fetch url " + url)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405")
	return http.DefaultClient.Do(req)
}

func Fetch() {

	db_feeds := (&postgres.RssFeedRepository{}).GetFeeds()
	for _, db_feed := range db_feeds {

		if db_feed.LastSyncTime + int64(db_feed.SyncInterval) < time.Now().Unix() {

			log.Println("fetch rss: " + db_feed.Link)
			f, err := rss.FetchByFunc(fetchFunc, db_feed.Link)

			if err != nil {
				log.Println(err)
				(&postgres.RssFeedRepository{}).SetFeedAsBroken(db_feed.Id)
			} else {
				(&postgres.RssFeedRepository{}).UpdateFeed(db_feed.Id, f, time.Now().Unix())

				for _, item := range f.Items {
					(&postgres.RssFeedRepository{}).InsertRssItem(db_feed.Id, item)
				}
			}
		}
	}
}
