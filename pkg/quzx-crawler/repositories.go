package quzx_crawler

import "github.com/SlyMarbo/rss"

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

