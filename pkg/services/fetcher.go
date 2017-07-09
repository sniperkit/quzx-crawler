package services

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/demas/cowl-go/pkg/logging"
)

func doEvery(d time.Duration, f func()) {

	for {
		f()
		timer := time.NewTimer(d)
		<-timer.C
	}
}

func heartBeat() {
	println("working...")
}

func FetchNews() {

	log.Println("fetching news")
	// wait 1 minute to start postgresql
	timer := time.NewTimer(time.Second * 40)
	<-timer.C

	syncInterval, err := strconv.Atoi(os.Getenv("SYNCINTERVAL"))
	if err != nil {
		logging.LogInfo("SYNCINTERVAL was not defined")
		panic(err)
	} else {

		go doEvery(time.Minute*time.Duration(syncInterval), (&StackOverflowService{}).Fetch)

		// каждые 30 минут спрашиваем 1000 самых удачных вопросов за последние 3 дня
		go doEvery(time.Minute*30, (&StackOverflowService{}).FetchVotedQuestions)
		doEvery(time.Minute*60, heartBeat)

		//(&RssFeedService{}).Fetch()
		//(&HackerNewsService{}).Fetch()
		//(&StackOverflowService{}).RemoveOldQuestions()
	}
}
