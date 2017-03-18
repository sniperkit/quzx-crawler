package main

import (

	_ "github.com/lib/pq"

	"time"
	"os"
	"strconv"
	"log"
	"github.com/demas/cowl-go/pkg/services"
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
			(&services.StackOverflowService{}).Fetch()
			(&services.RssFeedService{}).Fetch()
			(&services.HackerNewsService{}).Fetch()

			timer := time.NewTimer(time.Minute * time.Duration(syncInterval))
			<- timer.C
		}
	}
}
