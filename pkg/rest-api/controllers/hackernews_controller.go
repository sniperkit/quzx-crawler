package controllers

import (
	"github.com/demas/cowl-go/pkg/postgres"
	"net/http"
	"encoding/json"
)

func GetUnreadedHackerNews(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.HackerNewsRepository{}).GetUnreadedHackerNews()
}

func SetHackerNewsAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(PostData)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.HackerNewsRepository{}).SetHackerNewsAsReaded(bodyData.Id)
	}

	return bodyData, err
}

func SetHackerNewsAsReadedFromTime (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetHackerNewsAsReadedFromTimeStruct struct {
		FromTime int64 `json:"fromTime"`
	}

	bodyData := new(SetHackerNewsAsReadedFromTimeStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.HackerNewsRepository{}).SetHackerNewsAsReadedFromTime(bodyData.FromTime)
	}

	return bodyData, err
}

func SetAllHackerNewsAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	(&postgres.HackerNewsRepository{}).SetAllHackerNewsAsReaded()
	return ResultOk{"ok"}, nil
}
