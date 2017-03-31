package tst

import (
	"testing"
	"log"
	"os"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/demas/cowl-go/pkg/postgres"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"time"
	_ "github.com/lib/pq"
	"github.com/SlyMarbo/rss"
)

func prepare() {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"))

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`DELETE FROM Settings`)
	db.Exec(`DELETE FROM HackerNews`)
	db.Exec(`DELETE FROM RssItem`)
	db.Exec(`DELETE FROM RssFeed`)
	db.Exec(`DELETE FROM StackQuestions`)
	db.Exec(`INSERT INTO Settings(Name, Value) VALUES('one', 'one_value')`)
	db.Exec(`INSERT INTO HackerNews(Id, By, Score, Time, Title, Type, Url, Readed, Favorite)
			VALUES(1, 'by', 10, 123456, 'title', 'type', 'http://google.com', 0, 0)`)

	insertFeedQuery := `INSERT INTO RssFeed(Title, Description, Link, UpdateUrl, ImageTitle, ImageUrl, ImageHeight,
			     		        ImageWidth, LastSyncTime, Total, Unreaded, SyncInterval, AlternativeName,
					        RssType, ShowContent, ShowOrder, Folder, LimitFull, LimitHeadersOnly, Broken)
	      		     VALUES('feed_title', 'description', 'www.ya.ru', '', '', '', 0, 0, 0, 0, 0, 1000, 'alternative',
	      		            2, 0, 0, 'devel', 0, 0, 0)
	      		     RETURNING Id`
	var feedId int
	rows, err := db.Query(insertFeedQuery)
	if err != nil {
		log.Print(err)
	}
	if rows.Next() {
		rows.Scan(&feedId)
	}

	db.Exec(`INSERT INTO RssItem(FeedId, Title, Summary, Content, Link, Date, ItemId, Readed, Favorite)
			VALUES($1, 'item_title', 'item_summary', 'content', 'www.google.com', 0, '0', 0, 0)`, feedId)

	db.Exec(`INSERT INTO StackQuestions(Title, Link, QuestionId, Tags, Score, AnswerCount, ViewCount,
				                    UserId, UserReputation, UserDisplayName, UserProfileImage,
				                    Classification, Details, Readed, CreationDate, Favorite, Classified)
			VALUES('title', 'www.google.com', 10, 'linux,docker', 5, 10, 100, 1, 12, 'user', 'profile',
			        'docker', '', 0,  0, 0, 1)`)
}

func TestMain(m *testing.M) {

	prepare()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetValueFromSettings(t *testing.T) {

	value := (&postgres.SettingsRepository{}).GetSettings("one")
	if value != "one_value" {
		t.Error(fmt.Sprintf("Expected value of 'some_value', but it was %s instead", value))
	}
}

func TestGetMissingValueFromSettings(t *testing.T) {

	value := (&postgres.SettingsRepository{}).GetSettings("two")
	if value != "" {
		t.Error(fmt.Sprintf("Expected empty string, but it was %s instead", value))
	}
}

func TestSetValueForSettings(t *testing.T) {

	(&postgres.SettingsRepository{}).SetSettings("new_key", "new_value")
	value := (&postgres.SettingsRepository{}).GetSettings("new_key")
	if value != "new_value" {
		t.Error(fmt.Sprintf("Expected value of 'new_value', but it was %s instead", value))
	}
}

func TestSetDuplicateValueForSettings(t *testing.T) {

	(&postgres.SettingsRepository{}).SetSettings("one", "new_one_value")
	value := (&postgres.SettingsRepository{}).GetSettings("one")
	if value != "new_one_value" {
		t.Error(fmt.Sprintf("Expected value of 'new_one_value', but it was %s instead", value))
	}
}

func TestCheckThatHackerNewsExists(t *testing.T) {

	value:= (&postgres.HackerNewsRepository{}).NewsExists(1)
	if value != true {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestCheckThatHackerNewsDoesntExists(t *testing.T) {

	value:= (&postgres.HackerNewsRepository{}).NewsExists(2)
	if value != false {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestInsertHackerNews(t *testing.T) {

	var insertedId int64 = 999
	hn := quzx_crawler.HackerNews{ Id:insertedId, By:"by_demas", Score:20, Time:654321, Title:"some_title",
		Type:"some_type", Url:"www.yandex.ru", Readed:0, Favorite:0 }
	(&postgres.HackerNewsRepository{}).InsertNews(hn)

	value:= (&postgres.HackerNewsRepository{}).NewsExists(insertedId)
	if value != true {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestGetFeeds(t *testing.T) {

	result := (&postgres.RssFeedRepository{}).GetFeeds()
	if len(result) != 1 {
		t.Error(fmt.Sprintf("Expected value of 1, but it was %s instead", len(result)))
	}
}

func TestUpdateFeeds(t *testing.T) {

	feeds := (&postgres.RssFeedRepository{}).GetFeeds()
	if len(feeds) != 1 {
		t.Error(fmt.Sprintf("Expected value of 1, but it was %s instead", len(feeds)))
	}

	feed := feeds[0]
	if feed.LastSyncTime != 0 {
		t.Error(fmt.Sprintf("Expected value of 0, but it was %s instead", feed.LastSyncTime))
	}

	var i = rss.Image{"ititle", "iurl", 0, 0}
	var f = rss.Feed{"nick", "new_title", "new_description", "some",
	"url", &i, nil, nil, time.Now(), 0, nil}

	(&postgres.RssFeedRepository{}).UpdateFeed(feed.Id, &f, 20)

	feeds = (&postgres.RssFeedRepository{}).GetFeeds()
	feed = feeds[0]
	if feed.LastSyncTime != 20 {
		t.Error(fmt.Sprintf("Expected value of 20, but it was %s instead", feed.LastSyncTime))
	}
}

func TestSetFeedAsBroken(t *testing.T) {

	feeds := (&postgres.RssFeedRepository{}).GetFeeds()
	if len(feeds) != 1 {
		t.Error(fmt.Sprintf("Expected value of 1, but it was %s instead", len(feeds)))
	}

	feed := feeds[0]
	if feed.Broken != 0 {
		t.Error(fmt.Sprintf("Expected value of 0, but it was %s instead", feed.LastSyncTime))
	}

	(&postgres.RssFeedRepository{}).SetFeedAsBroken(feed.Id)

	feeds = (&postgres.RssFeedRepository{}).GetFeeds()
	feed = feeds[0]
	if feed.Broken != 1 {
		t.Error(fmt.Sprintf("Expected value of 0, but it was %s instead", feed.LastSyncTime))
	}
}

func TestInsertRssItem(t *testing.T) {

	feeds := (&postgres.RssFeedRepository{}).GetFeeds()
	if len(feeds) != 1 {
		t.Error(fmt.Sprintf("Expected value of 1, but it was %s instead", len(feeds)))
	}

	item := rss.Item{"title", "sum", "cont", "cat", "link", time.Now(),
	"10", nil, false}

	feed := feeds[0]
	(&postgres.RssFeedRepository{}).InsertRssItem(feed.Id, &item)
}

func TestInsertStackOverflowQuestions(t *testing.T) {

	var tags = []string{"linux, docker"}
	var user = quzx_crawler.SOUser{10, 1, "", 10, "",
	"", ""}
	var q1 = quzx_crawler.SOQuestion { tags,  user, true, 10,
		20, 10, 100, 100, 30, "link", "title"}
	var q2 = quzx_crawler.SOQuestion { tags,  user, true, 10,
					   20, 10, 100, 100, 30, "link2", "title2"}
	var questions = []quzx_crawler.SOQuestion{q1, q2}

	err := (&postgres.StackOverflowRepository{}).InsertSOQuestions(questions, "stackoverflow")
	if err != nil {
		t.Error("Error while inserting StackOverflow Question")
	}
}

func TestRemoveOldStackOverflowQuestions(t *testing.T) {

	err := (&postgres.StackOverflowRepository{}).RemoveOldQuestions(time.Now().Unix())
	if err != nil {
		t.Error("Error while removing old questions")
	}
}