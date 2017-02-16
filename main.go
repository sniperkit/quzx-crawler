package main

import (

	_ "github.com/lib/pq"

	"github.com/demas/cowl-go/stackoverflow"
	"github.com/demas/cowl-go/feed"
	"github.com/demas/cowl-go/reddit"

)

func main() {

	stackoverflow.Fetch()
	feed.Fetch()
	reddit.Fetch()
}
