package tst

import (
	"fmt"
	"testing"

	"github.com/sniperkit/quzx-crawler/pkg/postgres"
)

func TestGetValueFromSettings(t *testing.T) {

	value := (&postgres.SettingsRepository{}).GetSettings("one")
	if value != "one_value" {
		t.Error(fmt.Sprintf("Expected value of 'some_value', but it was %s instead", value))
	}
}

func TestGetMissingValueFromSettings(t *testing.T) {

	value := (&postgres.SettingsRepository{}).GetSettings("two")
	if value != "" {
		t.Error(fmt.Sprintf("Expected empty string, but it was %s instead", value))
	}
}

func TestSetValueForSettings(t *testing.T) {

	(&postgres.SettingsRepository{}).SetSettings("new_key", "new_value")
	value := (&postgres.SettingsRepository{}).GetSettings("new_key")
	if value != "new_value" {
		t.Error(fmt.Sprintf("Expected value of 'new_value', but it was %s instead", value))
	}
}

func TestSetDuplicateValueForSettings(t *testing.T) {

	(&postgres.SettingsRepository{}).SetSettings("one", "new_one_value")
	value := (&postgres.SettingsRepository{}).GetSettings("one")
	if value != "new_one_value" {
		t.Error(fmt.Sprintf("Expected value of 'new_one_value', but it was %s instead", value))
	}
}
