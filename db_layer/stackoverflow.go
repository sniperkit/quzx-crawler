package db_layer

import (
	"github.com/demas/cowl-go/stackoverflow"
	"github.com/demas/cowl-go/classificator"
	"strings"
	"log"
	"fmt"
)

func Insert_so_Questions(questions []stackoverflow.SOQuestion) {

	tx := db.MustBegin()
	for _, q := range questions {

		classification, details := classificator.Classify(q)

		var id int
		err := db.Get(&id, fmt.Sprintf("SELECT count(*) FROM StackQuestions WHERE QuestionId = '%d'", q.Question_id))
		if err != nil {
			log.Fatal(err)
		}

		if id == 0 {
			_, err = tx.Exec("INSERT INTO stackquestions(" +
				"title, link, questionid, tags, score, answercount, viewcount, userid, userreputation, " +
				"userdisplayname, userprofileimage, classification, details) " +
				"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
				q.Title, q.Link, q.Question_id, strings.Join(q.Tags[:], ","), q.Score, q.Answer_count, q.View_count,
				q.Owner.User_id, q.Owner.Reputation, q.Owner.Display_name, q.Owner.Profile_image,
				classification, details)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	tx.Commit()
}
