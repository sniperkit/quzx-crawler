package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sniperkit/quzx-crawler/pkg/rest-api/services"
)

func GetTwitterFavourites(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&services.TwitterService{}).GetFavoritesTwits(mux.Vars(r)["name"])
}

func SetTwitUnfavorite(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(PostData)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&services.TwitterService{}).DestroyFavorites(bodyData.Id)
	}

	return bodyData, err
}
