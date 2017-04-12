package tst

import (
	"testing"
	"github.com/demas/cowl-go/pkg/postgres"
	"fmt"
	"github.com/demas/cowl-services/pkg/quzx"
	"log"
	"github.com/SlyMarbo/rss"
	"time"
)

const FEED_LINK = "www.some-site.com"
const LAST_SYNC_TIME = 0

func deleteAllFeeds() {
	(&postgres.RssFeedRepository{}).DeleteAllRssFeeds()
}

func insertFeed() int {

	feed := quzx.RssFeed{ 0,
			      "",
			      "desc",
			      FEED_LINK,
			      "upd_url",
			      "img_title",
			      "img_url",
			      0,
			      0,
			      LAST_SYNC_TIME,
			      0,
			      0,
			      100,
			      "alt_name",
			      1,
			      1,
			      1,
			      "test",
			      100,
			      100,
			      0 }

	return (&postgres.RssFeedRepository{}).InsertRssFeed(&feed)
}

func insertRssItem(feedId int) int {

	item := rss.Item{"title",
		"sum",
		"cont",
		"cat",
		"link",
		time.Now(),
		"10",
		nil,
		false}

	return (&postgres.RssFeedRepository{}).InsertRssItem(feedId, &item)
}

func TestDeleteAllFeeds(t *testing.T) {

	deleteAllFeeds()
	result := (&postgres.RssFeedRepository{}).GetFeeds()
	if len(result) != 0 {
		t.Error(fmt.Sprintf("DeleteAllFeeds: expected value of 0, but it was %s instead", len(result)))
	}
}

func TestInsertRssFeed(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()

	if feedId == 0 {
		t.Error("InsertRssFeed: error returning id")
	}

	result := (&postgres.RssFeedRepository{}).GetFeeds()
	if len(result) != 1 {
		t.Error(fmt.Sprintf("Expected value of 1, but it was %s instead", len(result)))
	}
}

func TestGetFeeds(t *testing.T) {

	deleteAllFeeds()
	result := (&postgres.RssFeedRepository{}).GetFeeds()
	if len(result) != 0 {
		t.Error(fmt.Sprintf("Expected value of 0, but it was %s instead", len(result)))
	}

	insertFeed()
	result = (&postgres.RssFeedRepository{}).GetFeeds()
	if len(result) != 1 {
		t.Error(fmt.Sprintf("Expected value of 1, but it was %s instead", len(result)))
	}
}

func TestGetUnreadRssFeedByFeedId(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()
	result, err := (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		log.Println(err)
		t.Error("Cannot get feed by id")
	}

	if result.Link != FEED_LINK {
		t.Error("Get RSS feed by id is broken")
	}
}

func TestUpdateFeedBeforeSync(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()

	feed, err := (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Update RSS feed: something was broken")
	}

	if feed.LastSyncTime != 0 {
		t.Error(fmt.Sprintf("Expected value of 0, but it was %s instead", feed.LastSyncTime))
	}

	(&postgres.RssFeedRepository{}).UpdateFeedBeforeSync("new_title", "new_description", "", "",
	"", 0, 0, 20, feedId)

	feed, err = (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Update RSS feed: something was broken")
	}

	if feed.LastSyncTime != 20 {
		t.Error(fmt.Sprintf("Expected value of 20, but it was %s instead", feed.LastSyncTime))
	}
}

func TestUpdateFeedAfterSync(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()

	feed, err := (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Update RSS feed: something was broken")
	}

	if feed.LastSyncTime != 0 {
		t.Error(fmt.Sprintf("Expected value of 0, but it was %s instead", feed.LastSyncTime))
	}

	(&postgres.RssFeedRepository{}).UpdateFeedAfterSync("www.new-link.com", 100, "new name",
		1, 1, 0, "folder", 2000, feedId)

	feed, err = (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Update RSS feed: something was broken")
	}

	if feed.LastSyncTime != 100 {
		t.Error(fmt.Sprintf("Expected value of 100, but it was %s instead", feed.LastSyncTime))
	}
}

func TestSetFeedAsBroken(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()

	feed, err := (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Set feed as broken: something was broken")
	}

	if feed.Broken != 0 {
		t.Error(fmt.Sprintf("Expected value of 0, but it was %s instead", feed.LastSyncTime))
	}

	(&postgres.RssFeedRepository{}).SetFeedAsBroken(feed.Id)

	feed, err = (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Set feed as broken: something was broken")
	}

	if feed.Broken != 1 {
		t.Error(fmt.Sprintf("Expected value of 1, but it was %s instead", feed.LastSyncTime))
	}
}

func TestInsertAndGetRssItem(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()
	itemId := insertRssItem(feedId)

	result, err := (&postgres.RssFeedRepository{}).GetRssItemsByFeedId(feedId)
	if err != nil {
		t.Error("Insert and get rss items: something was broken")
	}

	if len(result) != 1 {
		t.Error(fmt.Sprintf("expected 1, but it was %s instead", len(result)))
	}

	_ ,err = (&postgres.RssFeedRepository{}).GetRssItemById(itemId)
	if err != nil {
		t.Error("Insert and get rss items: something was broken")
	}
}

func TestSetRssItemAsReaded(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()
	itemId := insertRssItem(feedId)

	item ,err := (&postgres.RssFeedRepository{}).GetRssItemById(itemId)
	if err != nil {
		t.Error("Set rss item as readed: something was broken")
	}

	if item.Readed != 0 {
		t.Error(fmt.Sprintf("Set rss item as readed: expected 0, but it was %s instead", item.Readed))
	}

	(&postgres.RssFeedRepository{}).SetRssItemAsReaded(itemId)

	item ,err = (&postgres.RssFeedRepository{}).GetRssItemById(itemId)
	if err != nil {
		t.Error("Set rss item as readed: something was broken")
	}

	if item.Readed != 1 {
		t.Error(fmt.Sprintf("Set rss item as readed: expected 1, but it was %s instead", item.Readed))
	}
}

func TestSetRssFeedAsReaded(t *testing.T) {

	deleteAllFeeds()
	feedId := insertFeed()
	itemId := insertRssItem(feedId)

	feed, err := (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Set feed as broken: something was broken")
	}

	if feed.Unreaded != 1 {
		t.Error(fmt.Sprintf("Set rss feed as readed: expected 1, but it was %s instead", feed.Unreaded))
	}

	item ,err := (&postgres.RssFeedRepository{}).GetRssItemById(itemId)
	if err != nil {
		t.Error("Set rss feed as readed: something was broken")
	}

	if item.Readed != 0 {
		t.Error(fmt.Sprintf("Set rss feed as readed: expected 0, but it was %s instead", item.Readed))
	}

	(&postgres.RssFeedRepository{}).SetRssFeedAsReaded(feedId)

	item ,err = (&postgres.RssFeedRepository{}).GetRssItemById(itemId)
	if err != nil {
		t.Error("Set rss feed as readed: something was broken")
	}

	if item.Readed != 1 {
		t.Error(fmt.Sprintf("Set rss feed as readed: expected 1, but it was %s instead", item.Readed))
	}

	feed, err = (&postgres.RssFeedRepository{}).GetRssFeedById(feedId)
	if err != nil {
		t.Error("Set feed as broken: something was broken")
	}

	if feed.Unreaded != 0 {
		t.Error(fmt.Sprintf("Set rss feed as readed: expected 0, but it was %s instead", feed.Unreaded))
	}
}

