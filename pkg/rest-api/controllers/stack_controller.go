package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/demas/cowl-go/pkg/rest-api/postgres"
	pg "github.com/demas/cowl-go/pkg/postgres"
	"encoding/json"
)

func GetStackTags(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&pg.StackOverflowRepository{}).GetStackTags()
}

func GetSecondTagsByClassification(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.StackService{}).GetSecondTagByClassification(mux.Vars(r)["classification"])
}

func GetStackQuestionsByClassification(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.StackService{}).GetStackQuestionsByClassification(mux.Vars(r)["classification"])
}

func GetStackQuestionsByClassificationAndDetails(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.StackService{}).GetStackQuestionsByClassificationAndDetails(mux.Vars(r)["classification"],
										      mux.Vars(r)["details"])
}

func SetStackQuestionAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetStackQuestionAsReadedStruct struct {
		QuestionId int `json:"questionid"`
	}

	bodyData := new(SetStackQuestionAsReadedStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.StackService{}).SetStackQuestionAsReaded(bodyData.QuestionId)
	}

	return bodyData, err
}

func SetStackQuestionsAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(UniversalPostStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.StackService{}).SetStackQuestionsAsReadedByClassification(bodyData.Tag)
	}

	return bodyData, err
}

func SetStackQuestionsAsReadedFromTime (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(UniversalPostStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.StackService{}).SetStackQuestionsAsReadedByClassificationFromTime(bodyData.Tag, bodyData.FromTime)
	}

	return bodyData, err
}

