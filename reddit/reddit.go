package reddit

import (
	"github.com/demas/cowl-go/db_layer"
	"github.com/SlyMarbo/rss"

	"log"
	"time"
)

func Fetch() {

	db_reddits := db_layer.GetReddits()
	for _, db_reddit := range db_reddits {

		if db_reddit.LastSyncTime + int64(db_reddit.SyncInterval) < time.Now().Unix() {

			log.Println("fetch reddit: " + db_reddit.Link)
			f, err := rss.Fetch(db_reddit.Link)

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

