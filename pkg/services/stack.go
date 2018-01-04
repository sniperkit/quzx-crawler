package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/demas/cowl-go/pkg/logging"
	"github.com/demas/cowl-go/pkg/postgres"
	"github.com/demas/cowl-go/pkg/quzx-crawler"
)

// represent an implementation of quzx_crawler.StackOverflowService
type StackOverflowService struct {
}

const soKeyEnvVariable = "SOKEY"
const maxSOPages = 50
const soBaseUrl = "https://api.stackexchange.com/2.2/questions?page=%d&pagesize=100&fromdate=%d&order=asc&sort=creation&site=%s%s"
const votesUrl = "https://api.stackexchange.com/2.2/questions?page=%d&pagesize=100&fromdate=%d&order=desc&sort=votes&&site=stackoverflow%s"
const removeOldQuestionsInterval = -7 * 24 * time.Hour

var soSites = [8]string{"stackoverflow", "security", "codereview", "softwareengineering", "ru.stackoverflow", "superuser",
	"unix", "serverfault"}

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
		logging.PostgreLog{}.LogInfo("(n) " + url)

		// fetch data
		res, err := http.Get(url)
		if err != nil {
			logging.PostgreLog{}.LogError(err.Error())
		}

		jsn, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			logging.PostgreLog{}.LogError(err.Error())
		}

		// decode
		var p quzx_crawler.SOResponse

		err = json.Unmarshal(jsn, &p)
		if err != nil {
			logging.PostgreLog{}.LogError(err.Error())
		} else {
			result = append(result, p.Items...)
		}

		has_more = p.Has_more
		page = page + 1
	}

	return result
}

// 1000 (100 * 10) самых удачных вопросов за последние 3 дня
func (s *StackOverflowService) getVotedQuestions() []quzx_crawler.SOQuestion {

	var result []quzx_crawler.SOQuestion

	var fromTime int64
	fromTime = time.Now().Add(-24 * time.Hour * 3).Unix()
	logging.PostgreLog{}.LogInfo("=== fetch voted messages")

	for page := 1; page <= 10; page ++ {

		url := fmt.Sprintf(votesUrl, page, fromTime, s.key())
		logging.PostgreLog{}.LogInfo("(v) " + url)

		// fetch data
		res, err := http.Get(url)
		if err != nil {
			logging.PostgreLog{}.LogError(err.Error())
		}

		jsn, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			logging.PostgreLog{}.LogError(err.Error())
		}

		// decode
		var p quzx_crawler.SOResponse

		err = json.Unmarshal(jsn, &p)
		if err != nil {
			logging.PostgreLog{}.LogError(err.Error())
		} else {
			result = append(result, p.Items...)
		}

	}

	return result
}

func (s *StackOverflowService) Fetch() {

	lastSyncTime := getLastSyncTime("lastStackSyncTime", 2000)
	currentTime := time.Now().Unix()

	for _, soSite := range soSites {

		soQuestions := s.getNewMassages(lastSyncTime, soSite)
		for _, question := range soQuestions {
			classified_question := Classify(question, soSite)
			(&postgres.StackOverflowRepository{}).InsertSOQuestion(&classified_question, soSite)
		}
	}

	(&postgres.SettingsRepository{}).SetSettings("lastStackSyncTime", strconv.FormatInt(currentTime, 10))
}

func (s *StackOverflowService) FetchVotedQuestions() {

	soQuestions := s.getVotedQuestions()
	for _, question := range soQuestions {
		(&postgres.StackOverflowRepository{}).UpdateSOQuestion(&question)
	}
}

func (s *StackOverflowService) RemoveOldQuestions() {

	(&postgres.StackOverflowRepository{}).RemoveOldQuestions(time.Now().Add(removeOldQuestionsInterval).Unix())
}
