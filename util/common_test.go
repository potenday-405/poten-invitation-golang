package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	ti, err := StringToTime("20241200")
	t.Log(time.Time{})
	t.Log(make(time.Time))
	t.Log(ti == time.Time{})
	assert.Error(t, err)
	ti, err = StringToTime("202412001111")
	t.Log(ti == time.Time{})
	assert.NoError(t, err)
	_, err = StringToTime("20241200111a")
	assert.Error(t, err)
}
