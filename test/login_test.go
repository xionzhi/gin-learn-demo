package test

import (
	login "apps/web/login"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpLogin(t *testing.T) {
	data := HttpLogin()

	target := login.WebResponse{}
	if err := json.Unmarshal(data, &target); err != nil {
		t.Error(target)
	}

	assert.Equal(t, 1, target.Code)
}
