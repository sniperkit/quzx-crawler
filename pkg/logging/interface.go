package logging

import (
	"github.com/sniperkit/quzx-crawler/pkg/quzx-crawler"
)

type LogInterface interface {
	InsertLogMessage(message quzx_crawler.LogMessage)
	LogInfo(message string)
	LogError(message string)
	LogMessage(message quzx_crawler.LogMessage)
}
