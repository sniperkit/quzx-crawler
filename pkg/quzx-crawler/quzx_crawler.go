package quzx_crawler

import (
	"database/sql"
	"github.com/SlyMarbo/rss"
)

type Settings struct {
	Name  string
	Value string
}

type RssFeed struct {
	Id               int
	Title            sql.NullString
	Description      sql.NullString
	Link             string
	UpdateUrl        sql.NullString
	ImageTitle       sql.NullString
	ImageUrl         sql.NullString
	ImageHeight      sql.NullInt64
	ImageWidth       sql.NullInt64
	LastSyncTime     int64
	Total            sql.NullInt64
	Unreaded         sql.NullInt64
	SyncInterval     int
	AlternativeName  string
	RssType          int
	ShowContent      int
	ShowOrder        int
	Folder           string
	LimitFull        int
	LimitHeadersOnly int
	Broken           int
}

type RssItem struct {
	Id      int
	FeedId  int
	Title   string
	Summary string
	Content string
	Link    string
	Date    int64
	ItemId  string
	Readed  int
}

type SOUser struct {
	Reputation    int
	User_id       int
	User_type     string
	Accept_rate   int
	Profile_image string
	Display_name  string
	Link          string
}

type SOQuestion struct {
	Tags               []string
	Owner              SOUser
	Is_answered        bool
	View_count         int
	Answer_count       int
	Score              int
	Last_activity_date uint32
	Creation_date      uint32
	Question_id        uint32
	Link               string
	Title              string
}

type SOResponse struct {
	Items           []SOQuestion
	Has_more        bool
	Quota_max       int
	Quota_remaining int
}

type HackerNews struct {
	Id     int64
	By     string
	Score  int
	Time   int64
	Title  string
	Type   string
	Url    string
	Readed int
}

type HackerNewsRepository interface {
	NewsExists(id int64) bool
	InsertNews(n HackerNews)
}

type RssFeedRepository interface {
	GetFeeds() []RssFeed
	UpdateFeed(id int, feed *rss.Feed, lastSyncTime int64)
	SetFeedAsBroken(id int)
	InsertRssItem(feed_id int, i *rss.Item)
}

type StackOverflowRepository interface {
	InsertSOQuestions(questions []SOQuestion, site string)
}

type SettingsRepository interface {
	GetSettings(key string) string
	SetSettings(key string, value string)
}
