package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sniperkit/quzx-crawler/pkg/postgres"
	"github.com/sniperkit/quzx-crawler/pkg/rest-api/quzx"
)

func GetUnreadRssFeeds(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	rss_type, _ := strconv.Atoi(mux.Vars(r)["rss_type"])
	return (&postgres.RssFeedRepository{}).GetUnreadRssFeeds(rss_type)
}

func GetAllRssFeeds(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.RssFeedRepository{}).GetFeeds(), nil
}

func GetRssFeedById(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	return (&postgres.RssFeedRepository{}).GetRssFeedById(id)
}

func GetRssItemsByFeedId(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	var news []*quzx.RssItem
	feed_id, err := strconv.Atoi(mux.Vars(r)["feed_id"])
	if err == nil {
		news, err = (&postgres.RssFeedRepository{}).GetRssItemsByFeedId(feed_id)
	}
	return news, err
}

func PutRssFeed(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	d := new(quzx.RssFeed)
	err := json.NewDecoder(r.Body).Decode(&d)
	if err == nil {
		(&postgres.RssFeedRepository{}).UpdateFeedAfterSync(d.Link, d.LastSyncTime, d.AlternativeName,
			d.RssType, d.ShowContent, d.ShowOrder, d.Folder, d.SyncInterval, d.Id)
	}

	return d, err
}

func PostRssFeed(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(quzx.RssFeed)
	err := json.NewDecoder(r.Body).Decode(&bodyData)

	if err == nil {
		(&postgres.RssFeedRepository{}).InsertRssFeed(bodyData)
	}

	return bodyData, err
}

func SetRssItemAsReaded(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(PostData)
	err := json.NewDecoder(r.Body).Decode(&bodyData)

	if err == nil {
		(&postgres.RssFeedRepository{}).SetRssItemAsReaded(int(bodyData.Id))
	}

	return bodyData, err
}

func SetRssFeedAsReaded(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetRssFeedAsReadedStruct struct {
		FeedId int `json:"feed_id"`
	}

	bodyData := new(SetRssFeedAsReadedStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)

	if err == nil {
		(&postgres.RssFeedRepository{}).SetRssFeedAsReaded(bodyData.FeedId)
	}

	return bodyData, err
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	feedid, err := strconv.Atoi(mux.Vars(r)["id"])
	(&postgres.RssFeedRepository{}).UnsubscribeRssFeed(feedid)
	return ResultOk{"ok"}, err
}
