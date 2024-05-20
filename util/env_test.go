package util

import (
	"os"
	"testing"
)

func TestEnvInitializer(t *testing.T) {
	if err := EnvInitializer(); err != nil {
		t.Error(err)
	}
	t.Log(os.Getenv("MYSQL_ACCOUNT"))
}
