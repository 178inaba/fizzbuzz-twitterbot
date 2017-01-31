package main

import (
	"fmt"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

type twitterToken struct {
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
}

func main() {
	tt, err := flagValidation()
	if err != nil {
		log.Fatalf("Flag error: %s.", err)
	}

	// Create twitter client.
	anaconda.SetConsumerKey(tt.consumerKey)
	anaconda.SetConsumerSecret(tt.consumerSecret)
	api := anaconda.NewTwitterApi(tt.accessToken, tt.accessTokenSecret)

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
		}

		log.WithField("id", t.Id).Infof("Success!")
		time.Sleep(time.Minute)
	}
}

func flagValidation() (*twitterToken, error) {
	ck := kingpin.Flag("consumer-key", "Twitter consumer key.").Envar("CONSUMER_KEY").String()
	cs := kingpin.Flag("consumer-secret", "Twitter consumer secret.").Envar("CONSUMER_SECRET").String()
	at := kingpin.Flag("access-token", "Twitter access token.").Envar("ACCESS_TOKEN").String()
	ats := kingpin.Flag("access-token-secret", "Twitter access token secret.").Envar("ACCESS_TOKEN_SECRET").String()
	kingpin.Parse()

	if *ck == "" {
		return nil, errors.New("consumer key is not set")
	}

	if *cs == "" {
		return nil, errors.New("consumer secret is not set")
	}

	if *at == "" {
		return nil, errors.New("access token is not set")
	}

	if *ats == "" {
		return nil, errors.New("access token secret is not set")
	}

	return &twitterToken{
		consumerKey:       *ck,
		consumerSecret:    *cs,
		accessToken:       *at,
		accessTokenSecret: *ats,
	}, nil
}

func tweetText(num uint64) string {
	tweet, isFB := fizzBuzz(num)
	if isFB {
		// Add number hashtag.
		tweet = fmt.Sprintf("%s #%d", tweet, num)
	}

	return tweet
}

func fizzBuzz(num uint64) (string, bool) {
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
