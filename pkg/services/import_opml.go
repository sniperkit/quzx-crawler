package services

import (
	"log"

	"github.com/demas/cowl-go/pkg/postgres"
	"github.com/demas/cowl-services/pkg/quzx"
	"github.com/gilliek/go-opml/opml"
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
		0}

	_, err := (&postgres.RssFeedRepository{}).GetRssFeedByUrl(url)
	if err != nil {

		log.Println("Import feed: " + url)
		(&postgres.RssFeedRepository{}).InsertRssFeed(&feed)
	}
}

func ImportOpml() {

	log.Println("Importing OPML file")

	doc, err := opml.NewOPMLFromFile("d:\\bazqux.xml")
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
