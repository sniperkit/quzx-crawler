package hackernews

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/demas/cowl-go/types"
	"encoding/json"
	"github.com/demas/cowl-go/db_layer"
	"time"
	"strconv"
)

func GetNews() {

	idsUrl := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	newsUrl := "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"

	var syncInterval int64 = 60 * 30
	var lastSyncTime int64
	var err error

	lastSyncTimeStr := db_layer.GetSettings("lastHackerNewsSyncTime")
	if lastSyncTimeStr == "" {
		lastSyncTime = time.Now().Unix() - syncInterval - 1
	} else {
		lastSyncTime, err = strconv.ParseInt(lastSyncTimeStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	currentTime := time.Now().Unix()

	if lastSyncTime + syncInterval > currentTime {
		return
	}

	res, err := http.Get(idsUrl)
	if err != nil {
		log.Fatal(err)
	}
	jsn, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// decode
	var ids []int64
	err = json.Unmarshal(jsn, &ids)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, id := range ids {

			if db_layer.NewsDoesntExists(id) {
				log.Println("hacker news: " + fmt.Sprintf(newsUrl, id))
				res, err := http.Get(fmt.Sprintf(newsUrl, id))
				if err != nil {
					log.Fatal(err)
				}
				jsn, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}

				// decode
				var news types.HackerNews
				err = json.Unmarshal(jsn, &news)
				if err != nil {
					log.Fatal(err)
				} else {
					db_layer.InsertNews(news)
				}
			}
		}
	}

	db_layer.SetSettings("lastHackerNewsSyncTime",  strconv.FormatInt(currentTime, 10))
}
