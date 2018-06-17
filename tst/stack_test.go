package tst

import (
	"fmt"
	"testing"
	"time"

	"github.com/sniperkit/quzx-crawler/pkg/postgres"
	"github.com/sniperkit/quzx-crawler/pkg/quzx-crawler"
)

const LINK = "link"

func deleteAllStackQuestions() {
	(&postgres.StackOverflowRepository{}).DeleteAllQuestions()
}

func deleteAllStackTags() {
	(&postgres.StackOverflowRepository{}).DeleteAllStackTags()
}

func insertStackOverflowQuestion() int {

	var tags = []string{"linux, docker"}
	var user = quzx_crawler.SOUser{10,
		1,
		"",
		10,
		"",
		"",
		""}
	var q = quzx_crawler.SOQuestion{tags,
		user,
		true,
		10,
		20,
		10,
		100,
		100,
		30,
		LINK,
		"title"}

	return (&postgres.StackOverflowRepository{}).InsertSOQuestion(&q, "stackoverflow")
}

func TestInsertAndGetStackOverflowQuestions(t *testing.T) {

	deleteAllStackQuestions()
	id := insertStackOverflowQuestion()

	q, err := (&postgres.StackOverflowRepository{}).GetStackQuestionById(id)

	if err != nil {
		t.Error("Insert and get stack question: something was broken")
	}

	if q.Link != LINK {
		t.Error("Insert and get stack question: incorrect value")
	}
}

func TestRemoveOldStackOverflowQuestions(t *testing.T) {

	deleteAllStackQuestions()
	insertStackOverflowQuestion()

	err := (&postgres.StackOverflowRepository{}).RemoveOldQuestions(time.Now().Unix())
	if err != nil {
		t.Error("Error while removing old questions")
	}
}

func TestInsertAndGetStackTag(t *testing.T) {

	deleteAllStackTags()

	var tag = quzx_crawler.StackTag{"lunux", 2, 0}
	(&postgres.StackOverflowRepository{}).InsertStackTag(&tag)

	tags, err := (&postgres.StackOverflowRepository{}).GetStackTags()
	if err != nil {
		t.Error("Insert and get stack tag: something was broken")
	}

	if len(tags) != 1 {
		t.Error(fmt.Sprintf("Insert and get stack tag: incorrect value: %s", len(tags)))
	}
}
