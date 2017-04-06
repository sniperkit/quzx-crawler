package logging

import (
	"log"
	"gopkg.in/mgo.v2"
	"time"
	"os"
)

type Message struct {
	Moment int64
	Application string
	Level string
	Message string
}

var mongo *mgo.Session

func init() {

	var err error
	mongo, err = mgo.Dial(os.Getenv("MONGODB"))
	if err != nil {
		log.Println(err)
	}
}

func LogInfo(message string) {

	LogMessage(Message{ Moment: time.Now().Unix(), Application: "crawler", Level: "info", Message: message })
}

func LogError(message string) {

	LogMessage(Message{ Moment: time.Now().Unix(), Application: "crawler", Level: "error", Message: message })
}

func LogMessage(message Message) {

	c := mongo.DB("quzx").C("logs")
	err := c.Insert(&message)
	if err != nil {
		log.Println(err)
	}

	if message.Level == "info" {
		log.Println(message.Message)
	} else if message.Level == "error" {
		log.Fatal(message.Message)
	}
}
