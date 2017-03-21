package postgres

import (
	"log"
	"strings"

	"github.com/demas/cowl-go/pkg/classificator"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

// represent a PostgreSQL implementation of quzx_crawler.StackOverflowRepository
type StackOverflowRepository struct {
}

// InsertSOQuestions : insert StackOverflow questions in database
func (r *StackOverflowRepository) InsertSOQuestions(questions []quzx_crawler.SOQuestion, site string) {

	tx := db.MustBegin()
	for _, q := range questions {

		classification, details := classificator.Classify(q, site)

		insertQuery := `INSERT INTO stackquestions(title, link, questionid, tags, score, answercount, viewcount, userid, userreputation, userdisplayname, 
		                                           userprofileimage, classification, details, creationdate, readed) 
						VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT DO NOTHING`

		_, err := tx.Exec(insertQuery,
			q.Title,
			q.Link,
			q.Question_id,
			strings.Join(q.Tags[:], ","),
			q.Score,
			q.Answer_count,
			q.View_count,
			q.Owner.User_id,
			q.Owner.Reputation,
			q.Owner.Display_name,
			q.Owner.Profile_image,
			classification,
			details,
			q.Creation_date,
			0)

		if err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()
}

func (r *StackOverflowRepository) RemoveOldQuestions(fromTime int64) {

	tx := db.MustBegin()
	deleteQuery := `DELETE FROM StackQuestions WHERE Classification = '' AND CreationDate < $1`
	_, err := tx.Exec(deleteQuery, fromTime)
	if err != nil {
		log.Println(err)
	}

	tx.Commit()
}
