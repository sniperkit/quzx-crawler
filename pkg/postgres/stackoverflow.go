package postgres

import (
	"log"
	"strings"

	"github.com/demas/cowl-go/pkg/logging"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"github.com/demas/cowl-services/pkg/quzx"
)

// represent a PostgreSQL implementation of quzx_crawler.StackOverflowRepository
type StackOverflowRepository struct {
}

func (s *StackOverflowRepository) GetStackQuestionById(id int) (*quzx.StackQuestion, error) {

	var item quzx.StackQuestion
	selectQuery := `SELECT Title, Link, QuestionId, Tags, CreationDate, Favorite, Classified
			FROM StackQuestions WHERE Id = $1`
	err := db.Get(&item, selectQuery, id)
	return &item, err
}

func (r *StackOverflowRepository) InsertSOQuestion(question *quzx_crawler.SOQuestion, site string) int {

	insertQuery := `INSERT INTO stackquestions
				(Title, Link, Questionid, Tags, Score, AnswerCount, ViewCount,
				 Userid, UserReputation, UserDisplayname, UserProfileImage, Classification,
				 Details, Creationdate, Readed, Favorite, Classified, Site)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
			ON CONFLICT DO NOTHING
			RETURNING Id`

	tx := db.MustBegin()

	rows, err := tx.Query(insertQuery,
		question.Title,
		question.Link,
		question.Question_id,
		strings.Join(question.Tags[:], ","),
		question.Score,
		question.Answer_count,
		question.View_count,
		question.Owner.User_id,
		question.Owner.Reputation,
		question.Owner.Display_name,
		question.Owner.Profile_image,
		"",
		"",
		question.Creation_date,
		0,
		0,
		0,
		site)

	if err != nil {
		logging.LogInfo(err.Error())
	}

	var id int = 0
	if rows.Next() {
		rows.Scan(&id)
	}

	tx.Commit()

	return id
}

func (r *StackOverflowRepository) UpdateSOQuestion(question *quzx_crawler.SOQuestion) {

	updateQuery := `UPDATE StackQuestions
				    SET Score = $1,
				        AnswerCount = $2,
				        ViewCount = $3
				    WHERE QuestionId = $4`

	tx := db.MustBegin()

	_, err := tx.Exec(updateQuery,
		question.Score,
		question.Answer_count,
		question.View_count,
		question.Question_id)

	if err != nil {
		logging.LogInfo(err.Error())
	}

	tx.Commit()
}

func (r *StackOverflowRepository) RemoveOldQuestions(fromTime int64) error {

	tx := db.MustBegin()

	deleteQuery := `DELETE FROM StackQuestions WHERE Classification = '' AND CreationDate < $1`
	_, err := tx.Exec(deleteQuery, fromTime)
	if err != nil {
		logging.LogInfo(err.Error())
	}

	deleteQuery = `DELETE FROM StackQuestions WHERE READED = 1`
	_, err = tx.Exec(deleteQuery)
	if err != nil {
		logging.LogInfo(err.Error())
	}

	tx.Commit()

	return err
}

func (r *StackOverflowRepository) DeleteAllQuestions() error {

	tx := db.MustBegin()

	deleteQuery := `DELETE FROM StackQuestions`
	_, err := tx.Exec(deleteQuery)
	if err != nil {
		logging.LogInfo(err.Error())
	}

	tx.Commit()

	return err
}

func (r *StackOverflowRepository) InsertStackTag(tag *quzx_crawler.StackTag) int {

	insertQuery := `INSERT INTO StackTags(Classification, Unreaded, Hidden)
			VALUES($1, $2, $3)
			ON CONFLICT DO NOTHING
			RETURNING Id`

	tx := db.MustBegin()

	rows, err := tx.Query(insertQuery, tag.Classification, tag.Unreaded, tag.Hidden)

	if err != nil {
		logging.LogInfo(err.Error())
	}

	var id int = 0
	if rows.Next() {
		rows.Scan(&id)
	}

	tx.Commit()

	return id
}

func (r *StackOverflowRepository) GetStackTags() ([]*quzx_crawler.StackTag, error) {

	selectQuery := `SELECT Classification, Unreaded
	                FROM StackTags
	                WHERE Unreaded > 0 and Hidden = 0`

	result := []*quzx_crawler.StackTag{}
	rows, err := db.Query(selectQuery)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := quzx_crawler.StackTag{}
			rows.Scan(&q.Classification, &q.Unreaded)
			result = append(result, &q)
		}
	}

	return result, err
}

func (r *StackOverflowRepository) DeleteAllStackTags() error {

	tx := db.MustBegin()

	deleteQuery := `DELETE FROM StackTags`
	_, err := tx.Exec(deleteQuery)
	if err != nil {
		logging.LogInfo(err.Error())
	}

	tx.Commit()

	return err
}
