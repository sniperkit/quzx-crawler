package quzxutil

import (
	"os"
)

func GetParameter(name string) (string) {

	result := os.Getenv(name)

	if result == "" {

		switch name {
		case "SYNCINTERVAL":
			result = "3"
		case "DBUSER":
			result = "root"
		case "DBPASS":
			result = "root"
		case "DBHOST":
			result = "192.168.1.71"
		case "DBPORT":
			result = "5432"
		case "DBNAME":
			result = "rss"
		case "PORT":
			result = "4000"
		case "USER":
			result = "demas"
		case "PASS":
			result = "root"
		}
	}

	return result
}
