package controller

import (
	"strconv"
	"testing"
)

func TestAtoI(t *testing.T) {
	atoi, err := strconv.Atoi("")
	if err != nil {
		t.Log(err)
	}
	t.Log(atoi)
}
