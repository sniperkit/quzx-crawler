package logging

import "github.com/demas/cowl-go/pkg/quzx-crawler"

type LogInterface interface {
	InsertLogMessage(message quzx_crawler.LogMessage)
	LogInfo(message string)
	LogError(message string)
	LogMessage(message quzx_crawler.LogMessage)
}
