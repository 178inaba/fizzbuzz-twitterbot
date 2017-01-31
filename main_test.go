package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagValidation(t *testing.T) {
	strPointer := func(str string) *string {
		return &str
	}

	tt, err := flagValidation()
	assert.Nil(t, tt)
	assert.Error(t, err)

	consumerKey = strPointer("consumer-key")
	tt, err = flagValidation()
	assert.Nil(t, tt)
	assert.Error(t, err)

	consumerSecret = strPointer("consumer-secret")
	tt, err = flagValidation()
	assert.Nil(t, tt)
	assert.Error(t, err)

	accessToken = strPointer("access-token")
	tt, err = flagValidation()
	assert.Nil(t, tt)
	assert.Error(t, err)

	accessTokenSecret = strPointer("access-token-secret")
	tt, err = flagValidation()
	assert.Equal(t, *consumerKey, tt.consumerKey)
	assert.Equal(t, *consumerSecret, tt.consumerSecret)
	assert.Equal(t, *accessToken, tt.accessToken)
	assert.Equal(t, *accessTokenSecret, tt.accessTokenSecret)
	assert.NoError(t, err)
}

func TestTweetText(t *testing.T) {
	tweet := tweetText(7)
	assert.Equal(t, "7", tweet)

	tweet = tweetText(18)
	assert.Equal(t, "Fizz #18", tweet)

	tweet = tweetText(20)
	assert.Equal(t, "Buzz #20", tweet)

	tweet = tweetText(45)
	assert.Equal(t, "FizzBuzz #45", tweet)

	tweet = tweetText(math.MaxUint64)
	assert.Equal(t, "FizzBuzz #18446744073709551615", tweet)
}

func TestFizzBuzz(t *testing.T) {
	fb, isFB := fizzBuzz(7)
	assert.Equal(t, "7", fb)
	assert.False(t, isFB)

	fb, isFB = fizzBuzz(18)
	assert.Equal(t, "Fizz", fb)
	assert.True(t, isFB)

	fb, isFB = fizzBuzz(20)
	assert.Equal(t, "Buzz", fb)
	assert.True(t, isFB)

	fb, isFB = fizzBuzz(45)
	assert.Equal(t, "FizzBuzz", fb)
	assert.True(t, isFB)

	fb, isFB = fizzBuzz(math.MaxUint64)
	assert.Equal(t, "FizzBuzz", fb)
	assert.True(t, isFB)
}
