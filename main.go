package main

import (

	_ "github.com/lib/pq"

	"time"
	"os"
	"strconv"
	"log"

	"github.com/demas/cowl-go/pkg/services/hackernews"
	"github.com/demas/cowl-go/pkg/services/stackoverflow"
	"github.com/demas/cowl-go/pkg/services/rss"
)

func main() {

	// wait 1 minute to start postgresql
	timer := time.NewTimer(time.Minute * 1)
	<- timer.C

	syncInterval, err := strconv.Atoi(os.Getenv("SYNCINTERVAL"))
	if err != nil {
		log.Println("SYNCINTERVAL was not defined")
		panic(err)
	} else {
		for {
			stackoverflow.Fetch()
			rssfeeds.Fetch()
			hackernews.GetNews()

			timer := time.NewTimer(time.Minute * time.Duration(syncInterval))
			<- timer.C
		}
	}
}
