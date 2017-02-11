package main

import (

	_ "github.com/lib/pq"
	"../cowl-go/db_layer"

	"time"
	"strconv"
	"log"
	"../cowl-go/stackoverflow"
)

func main() {

	var lastSyncTime int64
	var err error

	lastSyncTimeStr := db_layer.GetSettings("lastStackSyncTime")
	if lastSyncTimeStr == "" {
		lastSyncTime = time.Now().Unix() - 2000
	} else {
		lastSyncTime, err = strconv.ParseInt(lastSyncTimeStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	currentTime := time.Now().Unix()

	res := stackoverflow.GetNewMassages(lastSyncTime)
	db_layer.Insert_so_Questions(res)

	db_layer.SetSettings("lastStackSyncTime",  strconv.FormatInt(currentTime, 10))
}
