package main

import (
	_ "github.com/lib/pq"

	"github.com/demas/cowl-go/pkg/services"
)

func main() {

	services.FetchNews()
}
