package stackoverflow

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/demas/cowl-go/db_layer"
	"github.com/demas/cowl-go/types"
	"time"
	"strconv"

)

func key() string {

	envVar := "SOKEY"

	if os.Getenv(envVar) != "" {
		return fmt.Sprintf("&key=%s", os.Getenv(envVar))
	} else {
		return ""
	}
}

func getNewMassages(fromTime int64) []types.SOQuestion {

	var result []types.SOQuestion
	page := 1
	has_more := true

	for has_more {

		url := fmt.Sprintf(
			"https://api.stackexchange.com/2.2/questions?page=%d&pagesize=100&fromdate=%d&order=asc&sort=creation&site=stackoverflow%s",
			page, fromTime, key())

		fmt.Println(url)

		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		jsn, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		// decode
		var p types.SOResponse

		err = json.Unmarshal(jsn, &p)
		if err != nil {
			log.Fatal(err)
		} else {
			result = append(result, p.Items...)
		}
		has_more = p.Has_more
		page = page + 1
	}

	return result
}

func Fetch() {

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

	res := getNewMassages(lastSyncTime)
	db_layer.Insert_so_Questions(res)

	db_layer.SetSettings("lastStackSyncTime",  strconv.FormatInt(currentTime, 10))
}

