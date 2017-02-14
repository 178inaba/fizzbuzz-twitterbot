package post

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWaitNextZeroSec(t *testing.T) {
	next := nextZeroSec()
	assert.Equal(t, 0, time.Now().Add(next).Second())
}

func TestTweetText(t *testing.T) {
	tweet := tweetText(false, false, 7)
	assert.Equal(t, "7", tweet)

	tweet = tweetText(true, false, 18)
	assert.Equal(t, "Fizz #18", tweet)

	tweet = tweetText(false, true, 20)
	assert.Equal(t, "Buzz #20", tweet)

	tweet = tweetText(true, true, 45)
	assert.Equal(t, "FizzBuzz #45", tweet)

	tweet = tweetText(true, true, math.MaxUint64)
	assert.Equal(t, "FizzBuzz #18446744073709551615", tweet)
}

func TestFizzbuzzText(t *testing.T) {
	text := fizzbuzzText(false, false)
	assert.Equal(t, "", text)

	text = fizzbuzzText(true, false)
	assert.Equal(t, "Fizz", text)

	text = fizzbuzzText(false, true)
	assert.Equal(t, "Buzz", text)

	text = fizzbuzzText(true, true)
	assert.Equal(t, "FizzBuzz", text)
}

func TestFizzBuzz(t *testing.T) {
	isFizz, isBuzz := fizzbuzz(7)
	assert.False(t, isFizz)
	assert.False(t, isBuzz)

	isFizz, isBuzz = fizzbuzz(18)
	assert.True(t, isFizz)
	assert.False(t, isBuzz)

	isFizz, isBuzz = fizzbuzz(20)
	assert.False(t, isFizz)
	assert.True(t, isBuzz)

	isFizz, isBuzz = fizzbuzz(45)
	assert.True(t, isFizz)
	assert.True(t, isBuzz)

	isFizz, isBuzz = fizzbuzz(math.MaxUint64)
	assert.True(t, isFizz)
	assert.True(t, isBuzz)
}

func BenchmarkLinkingStringPlus(b *testing.B) {
	boolSlice := [][]bool{[]bool{false, false}, []bool{true, false}, []bool{false, true}, []bool{true, true}}
	var j int
	for i := 0; i < b.N; i++ {
		isFizz, isBuzz := boolSlice[j][0], boolSlice[j][1]
		var tweet string
		if isFizz {
			tweet = "Fizz"
		}

		if isBuzz {
			tweet += "Buzz"
		}

		b.Log(tweet)

		j++
		if j > 3 {
			j = 0
		}
	}
}

func BenchmarkLinkingStringByteAppend(b *testing.B) {
	boolSlice := [][]bool{[]bool{false, false}, []bool{true, false}, []bool{false, true}, []bool{true, true}}
	var j int
	for i := 0; i < b.N; i++ {
		isFizz, isBuzz := boolSlice[j][0], boolSlice[j][1]
		var tweetByte []byte
		if isFizz {
			tweetByte = append(tweetByte, "Fizz"...)
		}

		if isBuzz {
			tweetByte = append(tweetByte, "Buzz"...)
		}

		tweet := string(tweetByte)
		b.Log(tweet)

		j++
		if j > 3 {
			j = 0
		}
	}
}

func BenchmarkLinkingStringCapByteAppend(b *testing.B) {
	boolSlice := [][]bool{[]bool{false, false}, []bool{true, false}, []bool{false, true}, []bool{true, true}}
	var j int
	for i := 0; i < b.N; i++ {
		isFizz, isBuzz := boolSlice[j][0], boolSlice[j][1]
		tweetByte := make([]byte, 0, 8)
		if isFizz {
			tweetByte = append(tweetByte, "Fizz"...)
		}

		if isBuzz {
			tweetByte = append(tweetByte, "Buzz"...)
		}

		tweet := string(tweetByte)
		b.Log(tweet)

		j++
		if j > 3 {
			j = 0
		}
	}
}

func BenchmarkBoolTwo(b *testing.B) {
	boolSlice := [][]bool{[]bool{false, false}, []bool{true, false}, []bool{false, true}, []bool{true, true}}
	var j int
	for i := 0; i < b.N; i++ {
		isFizz, isBuzz := boolSlice[j][0], boolSlice[j][1]
		b.Log(isFizz || isBuzz)

		j++
		if j > 3 {
			j = 0
		}
	}
}

func BenchmarkBoolStringLen(b *testing.B) {
	stringSlice := []string{"", "Fizz", "Buzz", "FizzBuzz"}
	var j int
	for i := 0; i < b.N; i++ {
		tweet := stringSlice[j]
		b.Log(len(tweet) > 0)

		j++
		if j > 3 {
			j = 0
		}
	}
}

func tweetText2(isFizz, isBuzz bool, num uint64) string {
	text := fizzbuzzText(isFizz, isBuzz)
	if len(text) > 0 {
		// Add hashtag prefix.
		text += " #"
	}

	return fmt.Sprintf("%s%d", text, num)
}

func BenchmarkTweetText(b *testing.B) {
	boolSlice := [][]bool{[]bool{false, false}, []bool{true, false}, []bool{false, true}, []bool{true, true}}
	var j int
	for i := 0; i < b.N; i++ {
		isFizz, isBuzz := boolSlice[j][0], boolSlice[j][1]
		tweet := tweetText(isFizz, isBuzz, math.MaxUint64)
		b.Log(tweet)

		j++
		if j > 3 {
			j = 0
		}
	}
}

func BenchmarkTweetText2(b *testing.B) {
	boolSlice := [][]bool{[]bool{false, false}, []bool{true, false}, []bool{false, true}, []bool{true, true}}
	var j int
	for i := 0; i < b.N; i++ {
		isFizz, isBuzz := boolSlice[j][0], boolSlice[j][1]
		tweet := tweetText2(isFizz, isBuzz, math.MaxUint64)
		b.Log(tweet)

		j++
		if j > 3 {
			j = 0
		}
	}
}
