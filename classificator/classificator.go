package classificator

import "github.com/demas/cowl-go/pkg/quzx-crawler"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Classify(question quzx_crawler.SOQuestion, site string) (classification string, details string) {

	if site == "stackoverflow" {
		if contains(question.Tags, "azure") {
			return "azure", ""
		} else if contains(question.Tags, "postgresql") {
			return "postgresql", ""
		} else if contains(question.Tags, "clojure") {
			return "clojure", ""
		} else if contains(question.Tags, "go") {
			return "go", ""
		} else if contains(question.Tags, "angular2") {
			return "angular2", ""
		} else if contains(question.Tags, "git") {
			return "git", ""
		} else if contains(question.Tags, "docker") {
			return "docker", ""
		} else if contains(question.Tags, "typescript") {
			return "typescript", ""
		} else if contains(question.Tags, "python") {
			return "python", ""
		} else if contains(question.Tags, "sql-server") {
			return "ms sql server", ""
		} else {
			return "", ""
		}
	} else if site == "security" {
		return "information security", ""
	} else if site == "codereview" {
		return "code review", ""
	} else {
		return "", ""
	}
}

