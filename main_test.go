package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagValidation(t *testing.T) {
	strPointer := func(str string) *string {
		return &str
	}

	err := flagValidation()
	assert.Error(t, err)

	consumerKey = strPointer("consumer-key")
	err = flagValidation()
	assert.Error(t, err)

	consumerSecret = strPointer("consumer-secret")
	err = flagValidation()
	assert.Error(t, err)

	accessToken = strPointer("access-token")
	err = flagValidation()
	assert.Error(t, err)

	accessTokenSecret = strPointer("access-token-secret")
	err = flagValidation()
	assert.NoError(t, err)

	// Tear down.
	consumerKey = strPointer("")
	consumerSecret = strPointer("")
	accessToken = strPointer("")
	accessTokenSecret = strPointer("")
}
