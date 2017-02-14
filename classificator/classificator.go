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

func Classify(question types.SOQuestion) (classification string, details string) {

	if contains(question.Tags, "clojure") {
		return "clojure", ""
	} else if contains(question.Tags, "go") {
		return "go", ""
	} else if contains(question.Tags, "python") {
		return "python", ""
	} else if contains(question.Tags, "azure") {
		return "azure", ""
	} else {
		return "", ""
	}
}

