package main

import (

	_ "github.com/lib/pq"

	"github.com/demas/cowl-go/hackernews"
	"github.com/demas/cowl-go/stackoverflow"
	"github.com/demas/cowl-go/feed"
)

func main() {
	stackoverflow.Fetch()
	feed.Fetch()
	hackernews.GetNews()
}
