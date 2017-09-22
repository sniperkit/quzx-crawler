package services

import (
	"log"
	"strconv"
	"time"

	"github.com/demas/cowl-go/pkg/logging"
	"github.com/demas/cowl-go/pkg/quzxutil"
	"fmt"
)

func doEvery(d time.Duration, f func()) {

	for {
		f()
		timer := time.NewTimer(d)
		<-timer.C
	}
}

func heartBeat() {
	fmt.Println("working...")
}

func FetchNews() {

	log.Println("fetching news")
	// wait 1 minute to start postgresql
	timer := time.NewTimer(time.Second * 1)
	<-timer.C

	syncInterval, err := strconv.Atoi(quzxutil.GetParameter("SYNCINTERVAL"))
	if err != nil {
		logging.LogInfo("SYNCINTERVAL was not defined")
		panic(err)
	} else {

		go doEvery(time.Minute*time.Duration(syncInterval), (&StackOverflowService{}).Fetch)
		go doEvery(time.Minute*time.Duration(30), (&RssFeedService{}).Fetch)

		// каждые 30 минут спрашиваем 1000 самых удачных вопросов за последние 3 дня
		// go doEvery(time.Minute*30, (&StackOverflowService{}).FetchVotedQuestions)
		doEvery(time.Minute*60, heartBeat)


		//(&HackerNewsService{}).Fetch()
		//(&StackOverflowService{}).RemoveOldQuestions()
	}
}
