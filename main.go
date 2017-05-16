package main

import (
	_ "github.com/lib/pq"

	"log"
	"os"

	"github.com/demas/cowl-go/pkg/services"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("Specify command")
		return
	}

	if args[0] == "fetch" {
		services.FetchNews()
	} else if args[0] == "import-opml" {
		services.ImportOpml()
	} else {
		log.Println("unknown command")
	}

}
