package quzxutil

import (
	"os"
)

func GetParameter(name string) (string) {

	result := os.Getenv(name)

	if result == "" {

		switch name {
		case "SYNCINTERVAL":
			result = "10"
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
		case "TWICONKEY":
			result = "e7zN03mbqU0e2UbsrBDRo8iWj"
		case "TWICONSEC":
			result = "p38nB8qI3LypujakSHiaL1R8vQma0uuSjAdu7gu0D7AIU1pWws"
		case "TWIACCTOK":
			result = "402727952-Aq4uMyVH455DpdIdjpEKJBwb5NQXFxWYnqhcBRVX"
		case "TWIACCTOKSEC":
			result = "52naR3fWEOjbQqxk56So75gkXqfR0FcHTR95WyGb13Dsb"
		}
	}

	return result
}
