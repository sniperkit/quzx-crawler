package main

// go get github.com/alexflint/go-arg

import (
	_ "github.com/lib/pq"

	//"github.com/demas/cowl-go/pkg/services"
	"flag"
	"fmt"
	"os"
	"github.com/demas/cowl-go/pkg/services"
)

func main() {

	// parse command line arguments
	operation := flag.String("operation", "", "filename")
	filename := flag.String("filename", "", "filename")
	flag.Parse()
	fmt.Println(*operation)
	fmt.Println(*filename)

	switch *operation {

	case "import-opml":
		fmt.Println("Importing opml file")
		if len(*filename) == 0 {
			fmt.Println("Please provide filename")
			os.Exit(0)
		}
		services.ImportOpml(*filename)

	case "fetch-rss":
		fmt.Println("Fetch RSS")
		services.FetchNews()
	}
}
