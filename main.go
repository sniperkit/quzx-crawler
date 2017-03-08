package main

import (

	_ "github.com/lib/pq"

	"github.com/demas/cowl-go/hackernews"
	"github.com/demas/cowl-go/stackoverflow"
	"github.com/demas/cowl-go/feed"
	"time"
	"os"
	"strconv"
	"log"
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
			feed.Fetch()
			hackernews.GetNews()

			timer := time.NewTimer(time.Minute * time.Duration(syncInterval))
			<- timer.C
		}
	}
}
