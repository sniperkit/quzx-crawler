package services

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/sniperkit/quzx-crawler/pkg/postgres"
)

func ExportRssFeeds(filename string) {

	feeds := (&postgres.RssFeedRepository{}).GetFeeds()
	json, err := json.Marshal(feeds)
	if err != nil {
		log.Fatal("Cannot encode to JSON", err)
	}

	err = ioutil.WriteFile(filename, json, 0644)
	if err != nil {
		log.Fatal("Cannot write JSON to file")
	}
}
