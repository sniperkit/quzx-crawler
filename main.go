package main

import (

	_ "github.com/lib/pq"

	"time"
	"os"
	"strconv"
	"github.com/demas/cowl-go/pkg/services"
	"github.com/demas/cowl-go/pkg/logging"
)

func main() {

	// wait 1 minute to start postgresql
	timer := time.NewTimer(time.Minute * 1)
	<- timer.C

	syncInterval, err := strconv.Atoi(os.Getenv("SYNCINTERVAL"))
	if err != nil {
		logging.LogInfo("SYNCINTERVAL was not defined")
		panic(err)
	} else {
		for {
			(&services.StackOverflowService{}).Fetch()
			(&services.RssFeedService{}).Fetch()
			(&services.HackerNewsService{}).Fetch()

			(&services.StackOverflowService{}).RemoveOldQuestions()

			timer := time.NewTimer(time.Minute * time.Duration(syncInterval))
			<- timer.C
		}
	}
}
