package routers

import (
	"github.com/gorilla/mux"

	"github.com/sniperkit/quzx-crawler/pkg/rest-api/controllers"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	// stack (test)
	router.HandleFunc("/stack/tags", controllers.WrapHandler(controllers.GetStackTags))
	router.HandleFunc("/stack/secondtags/{classification}", controllers.WrapHandler(controllers.GetSecondTagsByClassification))
	router.HandleFunc("/stack/questions/{classification}", controllers.WrapHandler(controllers.GetStackQuestionsByClassification))
	router.HandleFunc("/stack/questions/{classification}/{details}", controllers.WrapHandler(controllers.GetStackQuestionsByClassificationAndDetails))
	router.HandleFunc("/stack/question-as-read", controllers.PostWrapHandler(controllers.SetStackQuestionAsReaded)).Methods("POST")
	router.HandleFunc("/stack/tags/as-read", controllers.PostWrapHandler(controllers.SetStackQuestionsAsReaded)).Methods("POST")
	router.HandleFunc("/stack/tags/from-time/as-read", controllers.PostWrapHandler(controllers.SetStackQuestionsAsReadedFromTime)).Methods("POST")

	// rss
	router.HandleFunc("/rss/unread/{rss_type}", controllers.WrapHandler(controllers.GetUnreadRssFeeds))
	router.HandleFunc("/rss/allfeeds", controllers.WrapHandler(controllers.GetAllRssFeeds))
	router.HandleFunc("/rss/{feed_id}/items", controllers.WrapHandler(controllers.GetRssItemsByFeedId))
	router.HandleFunc("/rss/item/as-read", controllers.PostWrapHandler(controllers.SetRssItemAsReaded)).Methods("POST")
	router.HandleFunc("/rss/feed/as-read", controllers.PostWrapHandler(controllers.SetRssFeedAsReaded)).Methods("POST")

	router.HandleFunc("/rss/feeds/{id}", controllers.WrapHandler(controllers.GetRssFeedById)).Methods("GET")
	router.HandleFunc("/rss/feeds", controllers.PostWrapHandler(controllers.PutRssFeed)).Methods("PUT")
	router.HandleFunc("/rss/feeds", controllers.PostWrapHandler(controllers.PostRssFeed)).Methods("POST")
	router.HandleFunc("/rss/feeds/{id}", controllers.PostWrapHandler(controllers.Unsubscribe)).Methods("DELETE")

	// twitter
	router.HandleFunc("/twitter/favorites/{name}", controllers.WrapHandler(controllers.GetTwitterFavourites))
	router.HandleFunc("/twitter/unfavorite", controllers.PostWrapHandler(controllers.SetTwitUnfavorite)).Methods("POST")

	// hacker news
	router.HandleFunc("/hn/unread", controllers.WrapHandler(controllers.GetUnreadedHackerNews))
	router.HandleFunc("/hn/as-read", controllers.PostWrapHandler(controllers.SetHackerNewsAsReaded)).Methods("POST")
	router.HandleFunc("/hn/all-as-read", controllers.PostWrapHandler(controllers.SetAllHackerNewsAsReaded)).Methods("POST")
	router.HandleFunc("/hn/fromtime-as-read", controllers.PostWrapHandler(controllers.SetHackerNewsAsReadedFromTime)).Methods("POST")

	// tags
	router.HandleFunc("/tags", controllers.WrapHandler(controllers.GetTags))
	router.HandleFunc("/tags/items/{tagId}", controllers.WrapHandler(controllers.GetTaggedItemsByTagId)).Methods("GET")
	router.HandleFunc("/tags/add-item", controllers.PostWrapHandler(controllers.InsertTaggedItem)).Methods("POST")
	router.HandleFunc("/tags/items/{id}", controllers.PostWrapHandler(controllers.DeleteTaggedItem)).Methods("DELETE")

	// bookmarks
	router.HandleFunc("/bookmarks", controllers.PostWrapHandler(controllers.PostBookmark)).Methods("POST")

	return router
}
