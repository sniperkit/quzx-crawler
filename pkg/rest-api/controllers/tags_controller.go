package controllers

import (
	"net/http"
	"github.com/demas/cowl-go/pkg/rest-api/postgres"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetTags(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.TagsService{}).GetTags()
}

func GetTaggedItemsByTagId(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	tagId, _ :=  strconv.Atoi(mux.Vars(r)["tagId"])
	return (&postgres.TagsService{}).GetTaggedItemsByTagId(tagId)
}

func InsertTaggedItem (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetTaggedItemStruct struct {
		ItemId int `json:"itemId"`
		TagId int `json:"tagId"`
		Source int `json:"source"`
	}

	bodyData := new(SetTaggedItemStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		if bodyData.Source == 1 {
			(&postgres.TagsService{}).InsertTaggedItemFromStockItem(bodyData.ItemId, bodyData.TagId)
		} else if bodyData.Source == 2 {
			(&postgres.TagsService{}).InsertTaggedItemFromRss(bodyData.ItemId, bodyData.TagId)
		}
	}

	return bodyData, err
}

func DeleteTaggedItem(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	id, _ :=  strconv.Atoi(mux.Vars(r)["id"])
	(&postgres.TagsService{}).DeleteTaggedItem(id)
	return ResultOk{"ok"}, nil
}

