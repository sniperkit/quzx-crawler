package reddit

import (
	"github.com/demas/cowl-go/db_layer"
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

	db_reddits := db_layer.GetReddits()
	for _, db_reddit := range db_reddits {

		if db_reddit.LastSyncTime + int64(db_reddit.SyncInterval) < time.Now().Unix() {

			log.Println("fetch reddit: " + db_reddit.Link)
			f, err := rss.FetchByFunc(fetchFunc, db_reddit.Link)

			if err != nil {
				log.Println(err)
			}
			db_layer.UpdateReddit(db_reddit.Id, f, time.Now().Unix())

			for _, item := range f.Items {
				db_layer.InsertRedditItem(db_reddit.Id, item)
			}
		}
	}
}

