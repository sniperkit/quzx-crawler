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
func (r *StackOverflowRepository) InsertSOQuestions(questions []quzx_crawler.SOQuestion, site string) error {

	tx := db.MustBegin()
	for _, q := range questions {

		classification, details := classificator.Classify(q, site)

		insertQuery := `INSERT INTO stackquestions
					(Title, Link, Questionid, Tags, Score, AnswerCount, ViewCount,
					 Userid, UserReputation, UserDisplayname, UserProfileImage, Classification,
					 Details, Creationdate, Readed, Favorite, Classified, Site)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
				ON CONFLICT DO NOTHING`

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
			0,
			0,
			0,
			site)

		if err != nil {
			log.Print(err)
			return err
		}
	}

	tx.Commit()
	return nil
}

func (r *StackOverflowRepository) RemoveOldQuestions(fromTime int64) error {

	tx := db.MustBegin()

	deleteQuery := `DELETE FROM StackQuestions WHERE Classification = '' AND CreationDate < $1`
	_, err := tx.Exec(deleteQuery, fromTime)
	if err != nil {
		log.Println(err)
	}

	deleteQuery = `DELETE FROM StackQuestions WHERE READED = 1`
	_, err = tx.Exec(deleteQuery)
	if err != nil {
		log.Println(err)
	}

	tx.Commit()

	return err
}
