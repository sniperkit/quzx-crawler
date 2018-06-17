package postgres

import (
	"log"

	"github.com/sniperkit/quzx-crawler/pkg/rest-api/quzx"
)

// represent a PostgreSQL implementation of quzx.StackService
type StackService struct {
}

func (s *StackService) GetSecondTagByClassification(classification string) (interface{}, error) {

	type Result struct {
		Details string `json:"details"`
		Count   int    `json:"count"`
	}

	result := []*Result{}
	selectQuery := `SELECT Details, COUNT(Id)
	                FROM StackQuestions
	                WHERE Classification = $1 and READED = 0
	                GROUP BY Details`

	err := db.Select(&result, selectQuery, classification)
	return result, err
}

func (s *StackService) GetStackQuestionsByClassification(classification string) ([]*quzx.StackQuestion, error) {

	log.Println("We are here !!!!!!!!!!!!!!!!!!!!!!!!")

	// TODO: обработка ошибок
	var result []*quzx.StackQuestion
	grm.Where("classification = ? and readed = 0", classification).Order("score desc").Find(result)
	return result, nil
}

func (s *StackService) GetStackQuestionsByClassificationAndDetails(classification string, details string) ([]*quzx.StackQuestion, error) {

	// TODO: обработка ошибок
	var result []*quzx.StackQuestion
	grm.Where("classification = ? and details = ? and readed = 0", classification, details).
		Order("score desc").Limit(15).Find(result)
	return result, nil
}

func (s *StackService) SetStackQuestionAsReaded(question_id int) {

	var question quzx.StackQuestion
	grm.Find(&question, question_id)
	question.Readed = 1
	grm.Save(&question)
}

func (s *StackService) SetStackQuestionsAsReadedByClassification(classification string) {

	grm.Model(quzx.StackQuestion{}).Where("classification = ?", classification).UpdateColumn("readed", 1)
}

func (s *StackService) SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64) {

	updateQuery := `UPDATE StackQuestions
	                SET READED = 1
	                WHERE Classification = $1 AND CreationDate < $2`

	tx := db.MustBegin()
	_, err := tx.Exec(updateQuery, classification, t)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}
