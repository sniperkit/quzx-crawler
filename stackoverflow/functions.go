package stackoverflow

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"os"
)

func key() string {

	envVar := "SOKEY"

	if os.Getenv(envVar) != "" {
		return fmt.Sprintf("&key=%s", os.Getenv(envVar))
	} else {
		return ""
	}
}

func GetNewMassages(fromTime int64) []SOQuestion {

	var result []SOQuestion
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
		var p SOResponse

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

