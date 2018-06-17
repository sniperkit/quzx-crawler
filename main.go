package main

// go get github.com/alexflint/go-arg

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/sniperkit/quzx-crawler/pkg/rest-api"
	"github.com/sniperkit/quzx-crawler/pkg/services"
	//"github.com/sniperkit/quzx-crawler/pkg/services"
)

func main() {

	// parse command line arguments
	operation := flag.String("operation", "", "filename")
	filename := flag.String("filename", "", "filename")

	flag.Parse()

	fmt.Println(*operation)

	switch *operation {

	case "import-opml":
		fmt.Println("Importing opml file")
		if len(*filename) == 0 {
			fmt.Println("Please provide filename")
			os.Exit(0)
		}
		services.ImportOpml(*filename)

	case "export-feeds":
		fmt.Println("Exporting feeds")
		if len(*filename) == 0 {
			fmt.Println("Please provide filename")
			os.Exit(0)
		}
		services.ExportRssFeeds(*filename)

	case "import-feeds":
		fmt.Println("Importing feeds")

	case "fetch-rss":
		fmt.Println("Fetch RSS")
		services.FetchNews()

	case "serve-rest-api":
		fmt.Println("Serve REST API")
		rest_api.Serve()
	}
}
