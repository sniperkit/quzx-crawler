package services

import (
	"time"
	"strconv"
	"github.com/demas/cowl-go/pkg/postgres"
)

func getLastSyncTime(key string, timeOffset int64) (int64, error) {

	lastSyncTimeStr := (&postgres.SettingsRepository{}).GetSettings(key)

	if lastSyncTimeStr == "" {
		return time.Now().Unix() - timeOffset, nil
	} else {
		return strconv.ParseInt(lastSyncTimeStr, 10, 64)
	}
}


