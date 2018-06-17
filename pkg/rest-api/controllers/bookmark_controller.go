package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sniperkit/quzx-crawler/pkg/rest-api/postgres"
	"github.com/sniperkit/quzx-crawler/pkg/rest-api/quzx"
)

/*	{
	"id": 5,
	"url": "1",
	"title":"title",
	"description":"description",
	"readitlater": 1,
	"tags": ["one", "two"]
} */

func PostBookmark(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(quzx.BookmarkPOST)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.BookmarkRepository{}).InsertBookmark(bodyData)
	}

	return bodyData, err
}
