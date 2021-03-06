package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sniperkit/quzx-crawler/pkg/logging"
	"github.com/sniperkit/quzx-crawler/pkg/postgres"
	"github.com/sniperkit/quzx-crawler/pkg/quzx-crawler"
)

// represent an implementation of quzx_crawler.HackerNewsService
type HackerNewsService struct {
}

const idsUrl = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
const newsUrl = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
const syncInterval = 30 * time.Minute

func (s *HackerNewsService) getMessagesIds() ([]int64, error) {

	res, err := http.Get(idsUrl)
	if err != nil {
		return nil, err
	}

	jsn, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	// decode
	var ids []int64
	err = json.Unmarshal(jsn, &ids)

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *HackerNewsService) fetchNews(id int64) (*quzx_crawler.HackerNews, error) {

	logging.PostgreLog{}.LogInfo("fetching from hacker news: " + fmt.Sprintf(newsUrl, id))
	res, err := http.Get(fmt.Sprintf(newsUrl, id))
	if err != nil {
		return nil, err
	}

	jsn, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	// decode
	var news quzx_crawler.HackerNews
	err = json.Unmarshal(jsn, &news)

	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (s *HackerNewsService) Fetch() {

	lastSyncTime := getLastSyncTime("lastHackerNewsSyncTime", int64(syncInterval.Seconds())+1)
	currentTime := time.Now().Unix()

	if lastSyncTime+int64(syncInterval.Seconds()) > currentTime {
		return
	}

	ids, err := s.getMessagesIds()
	if err != nil {
		logging.PostgreLog{}.LogError(err.Error())
	}

	for _, id := range ids {

		if !(&postgres.HackerNewsRepository{}).NewsExists(id) {

			news, err := s.fetchNews(id)
			if err != nil {
				logging.PostgreLog{}.LogInfo(err.Error())
			}
			(&postgres.HackerNewsRepository{}).InsertNews(*news)
		}
	}

	(&postgres.SettingsRepository{}).SetSettings("lastHackerNewsSyncTime", strconv.FormatInt(currentTime, 10))
}
