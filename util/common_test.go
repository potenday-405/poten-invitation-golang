package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToTime(t *testing.T) {
	_, err := StringToTime("20241200")
	assert.Error(t, err)
	_, err = StringToTime("202412001111")
	assert.NoError(t, err)
	_, err = StringToTime("20241200111a")
	assert.Error(t, err)
}
