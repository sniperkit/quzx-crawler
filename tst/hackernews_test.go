package tst

import (
	"testing"
	"fmt"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"github.com/demas/cowl-go/pkg/postgres"
)

const NEWS_ID = 999
const NEWS_TIME = 654321

func deleteAllHackerNews() {
	(&postgres.HackerNewsRepository{}).DeleteAllHackerNews()
}

func insertHackerNews() {

	hn := quzx_crawler.HackerNews{
		Id:NEWS_ID,
		By:"by_demas",
		Score:20,
		Time:NEWS_TIME,
		Title:"some_title",
		Type:"some_type",
		Url:"www.yandex.ru",
		Readed:0,
		Favorite:0 }

	(&postgres.HackerNewsRepository{}).InsertNews(hn)
}

func TestInsertHackerNews(t *testing.T) {

	deleteAllHackerNews()
	insertHackerNews()

	value:= (&postgres.HackerNewsRepository{}).NewsExists(NEWS_ID)
	if value != true {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestCheckThatHackerNewsExists(t *testing.T) {

	deleteAllHackerNews()
	insertHackerNews()

	value:= (&postgres.HackerNewsRepository{}).NewsExists(NEWS_ID)
	if value != true {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestCheckThatHackerNewsDoesntExists(t *testing.T) {

	deleteAllHackerNews()
	insertHackerNews()

	value := (&postgres.HackerNewsRepository{}).NewsExists(NEWS_ID + 1)
	if value != false {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestGetUnreadedHackerNews(t *testing.T) {

	deleteAllHackerNews()
	insertHackerNews()

	news, err := (&postgres.HackerNewsRepository{}).GetUnreadedHackerNews()
	if err != nil {
		t.Error("Get unreaded hacker news: something was broken")
	}

	if len(news) != 1 {
		t.Error("Get unreaded hacker news: incorrect count")
	}
}

func TestSetHackerNewsAsReaded(t *testing.T) {

	deleteAllHackerNews()
	insertHackerNews()

	(&postgres.HackerNewsRepository{}).SetHackerNewsAsReaded(NEWS_ID)

	news, err := (&postgres.HackerNewsRepository{}).GetHackerNewsById(NEWS_ID)
	if err != nil {
		t.Error("Set hacker news as readed: something was broken")
	}

	if news.Readed != 1 {
		t.Error("Set hacker news as readed: incorrect value")
	}
}

func TestSetHackerNewsAsReadedFromTime(t *testing.T) {

	deleteAllHackerNews()
	insertHackerNews()

	(&postgres.HackerNewsRepository{}).SetHackerNewsAsReadedFromTime(NEWS_TIME - 1)

	news, err := (&postgres.HackerNewsRepository{}).GetHackerNewsById(NEWS_ID)
	if err != nil {
		t.Error("Set hacker news as readed from time: something was broken")
	}

	if news.Readed != 0 {
		t.Error("Set hacker news as readed from time: incorrect value")
	}

	(&postgres.HackerNewsRepository{}).SetHackerNewsAsReadedFromTime(NEWS_TIME + 1)

	news, err = (&postgres.HackerNewsRepository{}).GetHackerNewsById(NEWS_ID)
	if err != nil {
		t.Error("Set hacker news as readed from time: something was broken")
	}

	if news.Readed != 1 {
		t.Error("Set hacker news as readed from time: incorrect value")
	}
}

func TestAllHAckerNewsAsReaded(t *testing.T) {

	deleteAllHackerNews()
	insertHackerNews()
	(&postgres.HackerNewsRepository{}).SetAllHackerNewsAsReaded()

	news, err := (&postgres.HackerNewsRepository{}).GetHackerNewsById(NEWS_ID)
	if err != nil {
		t.Error("Set hacker news as readed from time: something was broken")
	}

	if news.Readed != 1 {
		t.Error("Set hacker news as readed from time: incorrect value")
	}
}




