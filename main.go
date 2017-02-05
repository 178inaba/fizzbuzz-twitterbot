package main

import (
	"fmt"
	"os"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

var (
	consumerKey       = kingpin.Flag("consumer-key", "Twitter consumer key.").Envar("CONSUMER_KEY").String()
	consumerSecret    = kingpin.Flag("consumer-secret", "Twitter consumer secret.").Envar("CONSUMER_SECRET").String()
	accessToken       = kingpin.Flag("access-token", "Twitter access token.").Envar("ACCESS_TOKEN").String()
	accessTokenSecret = kingpin.Flag("access-token-secret", "Twitter access token secret.").Envar("ACCESS_TOKEN_SECRET").String()
)

type twitterToken struct {
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
}

func main() {
	kingpin.Parse()
	os.Exit(run())
}

func run() int {
	tt, err := flagValidation()
	if err != nil {
		log.Errorf("Flag error: %s.", err)
		return 1
	}

	// Create twitter client.
	anaconda.SetConsumerKey(tt.consumerKey)
	anaconda.SetConsumerSecret(tt.consumerSecret)
	api := anaconda.NewTwitterApi(tt.accessToken, tt.accessTokenSecret)

	// Start to 00 second.
	waitNextZeroSec()

	for i := uint64(1); ; i++ {
		tweet := tweetText(i)
		log.Infof("Tweet: %s", tweet)

		var t anaconda.Tweet
		for {
			var err error
			t, err = api.PostTweet(tweet, nil)
			if err == nil {
				break
			}

			log.Error(err)
			time.Sleep(time.Second)
		}

		log.WithField("id", t.Id).Infof("Success!")

		// Next post to 00 second.
		waitNextZeroSec()
	}
}

func nextZeroSec() time.Duration {
	n := time.Now()
	start := time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute()+1, 0, 0, time.Local)

	return start.Sub(n)
}

func waitNextZeroSec() {
	time.Sleep(nextZeroSec())
}

func flagValidation() (*twitterToken, error) {
	if *consumerKey == "" {
		return nil, errors.New("consumer key is not set")
	}

	if *consumerSecret == "" {
		return nil, errors.New("consumer secret is not set")
	}

	if *accessToken == "" {
		return nil, errors.New("access token is not set")
	}

	if *accessTokenSecret == "" {
		return nil, errors.New("access token secret is not set")
	}

	return &twitterToken{
		consumerKey:       *consumerKey,
		consumerSecret:    *consumerSecret,
		accessToken:       *accessToken,
		accessTokenSecret: *accessTokenSecret,
	}, nil
}

func tweetText(num uint64) string {
	tweet, isFB := fizzbuzz(num)
	if isFB {
		// Add number hashtag.
		tweet = fmt.Sprintf("%s #%d", tweet, num)
	}

	return tweet
}

func fizzbuzz(num uint64) (string, bool) {
	var fb string
	if num%3 == 0 {
		fb = "Fizz"
	}

	if num%5 == 0 {
		fb += "Buzz"
	}

	isFB := true
	if len(fb) == 0 {
		fb = fmt.Sprint(num)
		isFB = false
	}

	return fb, isFB
}
