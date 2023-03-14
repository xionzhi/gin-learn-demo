package test

import (
	login "apps/web/login"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	ret := Sum(2, 7)

	assert.Equal(t, 9, ret)
}

func TestHttpIndex(t *testing.T) {
	data := HttpIndex()

	target := login.WebResponse{}
	if err := json.Unmarshal(data, &target); err != nil {
		t.Error(target)
	}

	assert.Equal(t, 1, target.Code)
	assert.Equal(t, "hello world!", target.Message)
}
