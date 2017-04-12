package tst

import (
	"testing"
	"fmt"

	"github.com/demas/cowl-go/pkg/postgres"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"time"
	_ "github.com/lib/pq"
)

func TestCheckThatHackerNewsExists(t *testing.T) {

	value:= (&postgres.HackerNewsRepository{}).NewsExists(1)
	if value != true {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestCheckThatHackerNewsDoesntExists(t *testing.T) {

	value:= (&postgres.HackerNewsRepository{}).NewsExists(2)
	if value != false {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}

func TestInsertHackerNews(t *testing.T) {

	var insertedId int64 = 999
	hn := quzx_crawler.HackerNews{ Id:insertedId, By:"by_demas", Score:20, Time:654321, Title:"some_title",
		Type:"some_type", Url:"www.yandex.ru", Readed:0, Favorite:0 }
	(&postgres.HackerNewsRepository{}).InsertNews(hn)

	value:= (&postgres.HackerNewsRepository{}).NewsExists(insertedId)
	if value != true {
		t.Error(fmt.Sprintf("Expected value of 'true', but it was %s instead", value))
	}
}


func TestInsertStackOverflowQuestions(t *testing.T) {

	var tags = []string{"linux, docker"}
	var user = quzx_crawler.SOUser{10, 1, "", 10, "",
	"", ""}
	var q1 = quzx_crawler.SOQuestion { tags,  user, true, 10,
		20, 10, 100, 100, 30, "link", "title"}
	var q2 = quzx_crawler.SOQuestion { tags,  user, true, 10,
					   20, 10, 100, 100, 30, "link2", "title2"}
	var questions = []quzx_crawler.SOQuestion{q1, q2}

	err := (&postgres.StackOverflowRepository{}).InsertSOQuestions(questions, "stackoverflow")
	if err != nil {
		t.Error("Error while inserting StackOverflow Question")
	}
}

func TestRemoveOldStackOverflowQuestions(t *testing.T) {

	err := (&postgres.StackOverflowRepository{}).RemoveOldQuestions(time.Now().Unix())
	if err != nil {
		t.Error("Error while removing old questions")
	}
}