package tst

import (
	"testing"

	"github.com/demas/cowl-go/pkg/postgres"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"time"
	_ "github.com/lib/pq"
)


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