package quzx

import (
	"github.com/ChimeraCoder/anaconda"

	"github.com/sniperkit/quzx-crawler/pkg/quzx-crawler"
)

type FeedService interface {
	GetAllRssFeeds() ([]*RssFeed, error)
	GetUnreadRssFeeds(rssType int) ([]*RssFeed, error)
	GetRssFeedById(id int) (RssFeed, error)
	UpdateRssFeed(feed *RssFeed)
	InsertRssFeed(feed *RssFeed)
	GetRssItemsByFeedId(feed_id int) ([]*RssItem, error)
	SetRssItemAsReaded(id int)
	SetRssFeedAsReaded(feedId int)
	UnsubscribeRssFeed(feedId int)
}

type HackerNewsService interface {
	GetUnreadedHackerNews() ([]*HackerNews, error)
	SetHackerNewsAsReaded(id int64)
	SetHackerNewsAsReadedFromTime(t int64)
	SetAllHackerNewsAsReaded()
}

type StackService interface {
	GetStackTags() ([]*quzx_crawler.StackTag, error)
	GetStackQuestionById(id int) (*StackQuestion, error)
	GetStackQuestionsByClassification(classification string) ([]*StackQuestion, error)
	SetStackQuestionAsReaded(question_id int)
	SetStackQuestionsAsReadedByClassification(classification string)
	SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64)
}

type TagsService interface {
	GetTags() ([]*Tag, error)
	GetTaggedItemsByTagId(tagId int) ([]*TaggedItem, error)
	InsertTaggedItemFromStockItem(questionId int, tagId int)
	InsertTaggedItemFromRss(rssItemId int, tagId int)
	DeleteTaggedItem(id int)
}

type TwitterService interface {
	GetFavoritesTwits(name string) ([]anaconda.Tweet, error)
	DestroyFavorites(id int64)
}

type BookmarkRepository interface {
	InsertBookmark(bookmark Bookmark, tags []string)
}
