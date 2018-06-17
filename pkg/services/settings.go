package services

import (
	"strconv"
	"time"

	"github.com/sniperkit/quzx-crawler/pkg/postgres"
)

func getLastSyncTime(key string, timeOffset int64) int64 {

	lastSyncTimeStr := (&postgres.SettingsRepository{}).GetSettings(key)
	if lastSyncTimeStr == "" {
		return time.Now().Unix() - timeOffset
	}

	result, err := strconv.ParseInt(lastSyncTimeStr, 10, 64)
	if err != nil {
		return time.Now().Unix() - timeOffset
	}

	return result
}
