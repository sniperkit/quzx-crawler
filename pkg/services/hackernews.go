package services

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"encoding/json"
	"github.com/demas/cowl-go/pkg/postgres"
	"time"
	"strconv"
)

// represent an implementation of quzx_crawler.HackerNewsService
type HackerNewsService struct {
}

const idsUrl = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
const newsUrl = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
const syncInterval = 60 * 30 // 30 minutes

func (s *HackerNewsService) Fetch() {

	var lastSyncTime int64
	var err error

	lastSyncTimeStr := (&postgres.SettingsRepository{}).GetSettings("lastHackerNewsSyncTime")
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

	// get list of messages
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

			if !(&postgres.HackerNewsRepository{}).NewsExists(id) {

				// fetch each news
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
				var news quzx_crawler.HackerNews
				err = json.Unmarshal(jsn, &news)

				if err != nil {
					log.Fatal(err)
				} else {
					(&postgres.HackerNewsRepository{}).InsertNews(news)
				}
			}
		}
	}

	(&postgres.SettingsRepository{}).SetSettings("lastHackerNewsSyncTime",  strconv.FormatInt(currentTime, 10))
}
