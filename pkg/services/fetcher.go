package services

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/demas/cowl-go/pkg/logging"
)

func FetchNews() {

	log.Println("fetching news")
	// wait 1 minute to start postgresql
	// timer := time.NewTimer(time.Minute * 1)
	// <-timer.C

	syncInterval, err := strconv.Atoi(os.Getenv("SYNCINTERVAL"))
	if err != nil {
		logging.LogInfo("SYNCINTERVAL was not defined")
		panic(err)
	} else {
		for {
			(&StackOverflowService{}).Fetch()
			//(&RssFeedService{}).Fetch()
			//(&HackerNewsService{}).Fetch()

			//(&StackOverflowService{}).RemoveOldQuestions()

			timer := time.NewTimer(time.Minute * time.Duration(syncInterval))
			<-timer.C
		}
	}
}
