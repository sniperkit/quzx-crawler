package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/demas/cowl-go/pkg/postgres"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

// represent an implementation of quzx_crawler.StackOverflowService
type StackOverflowService struct {
}

const soKeyEnvVariable = "SOKEY"
const maxSOPages = 50
const soBaseUrl = "https://api.stackexchange.com/2.2/questions?page=%d&pagesize=100&fromdate=%d&order=asc&sort=creation&site=%s%s"
const removeOldQuestionsInterval = -7 * 24 * time.Hour

var   soSites = [3]string{"stackoverflow", "security", "codereview"}

func (s *StackOverflowService) key() string {

	if os.Getenv(soKeyEnvVariable) != "" {
		return fmt.Sprintf("&key=%s", os.Getenv(soKeyEnvVariable))
	}

	return ""
}

func (s *StackOverflowService) getNewMassages(fromTime int64, site string) []quzx_crawler.SOQuestion {

	var result []quzx_crawler.SOQuestion
	page := 1
	has_more := true

	for has_more && page <= maxSOPages {

		url := fmt.Sprintf(soBaseUrl, page, fromTime, site, s.key())
		log.Println(url)

		// fetch data
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
		var p quzx_crawler.SOResponse

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

func (s *StackOverflowService) Fetch() {

	var lastSyncTime int64
	var err error

	lastSyncTimeStr := (&postgres.SettingsRepository{}).GetSettings("lastStackSyncTime")
	if lastSyncTimeStr == "" {
		lastSyncTime = time.Now().Unix() - 2000
	} else {
		lastSyncTime, err = strconv.ParseInt(lastSyncTimeStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	currentTime := time.Now().Unix()

	for _, soSite := range soSites {
		soQuestions := s.getNewMassages(lastSyncTime, soSite)
		(&postgres.StackOverflowRepository{}).InsertSOQuestions(soQuestions, soSite)
	}

	(&postgres.SettingsRepository{}).SetSettings("lastStackSyncTime", strconv.FormatInt(currentTime, 10))
}

func (s *StackOverflowService) RemoveOldQuestions() {

	(&postgres.StackOverflowRepository{}).RemoveOldQuestions(time.Now().Add(removeOldQuestionsInterval).Unix())
}
