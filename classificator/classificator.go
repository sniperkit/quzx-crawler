package classificator

import "github.com/demas/cowl-go/types"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Classify(question types.SOQuestion, site string) (classification string, details string) {

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
		} else if contains(question.Tags, "python") {
			return "python", ""
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

