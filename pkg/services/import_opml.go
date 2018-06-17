package services

import (
	"log"

	"github.com/gilliek/go-opml/opml"

	"github.com/sniperkit/quzx-crawler/pkg/postgres"
	"github.com/sniperkit/quzx-crawler/pkg/rest-api/quzx"
)

func insertRssFeed(url string) {

	feed := quzx.RssFeed{0,
		"",
		"just_import",
		url,
		"",
		"",
		"",
		0,
		0,
		0,
		0,
		0,
		21600,
		"",
		1,
		1,
		1,
		"import",
		100,
		100,
		0,
		""}

	_, err := (&postgres.RssFeedRepository{}).GetRssFeedByUrl(url)
	if err != nil {

		log.Println("Import feed: " + url)
		(&postgres.RssFeedRepository{}).InsertRssFeed(&feed)
	}
}

func ImportOpml(filename string) {

	log.Println("Importing OPML file")

	doc, err := opml.NewOPMLFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, outline := range doc.Body.Outlines {

		if len(outline.Outlines) == 0 {
			insertRssFeed(outline.XMLURL)
		} else {
			for _, outline := range outline.Outlines {
				insertRssFeed(outline.XMLURL)
			}
		}
	}
}
