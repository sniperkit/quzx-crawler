package quzx_crawler

type RssFeedService interface {
	Fetch()
}

type StackOverflowService interface {
	Fetch()
	RemoveOldQuestions()
}

type HackerNewsService interface {
	Fetch()
}
