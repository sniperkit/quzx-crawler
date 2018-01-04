package services

import (
	"github.com/SlyMarbo/rss"
	"github.com/demas/cowl-go/pkg/postgres"

	"net/http"
	"time"

	"github.com/demas/cowl-go/pkg/logging"
)

// represent an implementation of quzx_crawler.RssFeedService
type RssFeedService struct {
}

const userAgent = "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"

func (s *RssFeedService) fetchFunc(url string) (resp *http.Response, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.PostgreLog{}.LogInfo("reddit: cannot fetch url " + url)
	}

	req.Header.Set("User-Agent", userAgent)
	return http.DefaultClient.Do(req)
}

func (s *RssFeedService) Fetch() {

	db_feeds := (&postgres.RssFeedRepository{}).GetFeeds()
	for _, db_feed := range db_feeds {

		if db_feed.LastSyncTime+int64(db_feed.SyncInterval) < time.Now().Unix() {

			logging.PostgreLog{}.LogInfo("fetch rss: " + db_feed.Link)
			f, err := rss.FetchByFunc(s.fetchFunc, db_feed.Link)

			if err != nil {
				logging.PostgreLog{}.LogInfo(err.Error())
				(&postgres.RssFeedRepository{}).SetFeedAsBroken(db_feed.Id, err.Error())
			} else {
				(&postgres.RssFeedRepository{}).UpdateFeedBeforeSync(f.Title, f.Description, f.UpdateURL,
					f.Image.Title, f.Image.URL, f.Image.Height, f.Image.Width, time.Now().Unix(), db_feed.Id)

				for _, item := range f.Items {
					(&postgres.RssFeedRepository{}).InsertRssItem(db_feed.Id, item)
				}
			}
		}
	}
}
