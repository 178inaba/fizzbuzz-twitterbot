package main

import (
	"fmt"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
)

var (
	consumerKey       = kingpin.Flag("consumer-key", "Twitter consumer key.").Envar("CONSUMER_KEY").String()
	consumerSecret    = kingpin.Flag("consumer-secret", "Twitter consumer secret.").Envar("CONSUMER_SECRET").String()
	accessToken       = kingpin.Flag("access-token", "Twitter access token.").Envar("ACCESS_TOKEN").String()
	accessTokenSecret = kingpin.Flag("access-token-secret", "Twitter access token secret.").Envar("ACCESS_TOKEN_SECRET").String()
)

func main() {
	kingpin.Parse()
	if *consumerKey == "" {
		log.Error("Consumer key is not set.")
		return
	}

	if *consumerSecret == "" {
		log.Error("Consumer secret is not set.")
		return
	}

	if *accessToken == "" {
		log.Error("Access token is not set.")
		return
	}

	if *accessTokenSecret == "" {
		log.Error("Access token secret is not set.")
		return
	}

	// Twitter.
	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	api := anaconda.NewTwitterApi(*accessToken, *accessTokenSecret)

	var i uint64
	for ; ; i++ {
		tweet := tweetText(i)
		log.Infof("Tweet: %s", tweet)

		var t anaconda.Tweet
		for {
			var err error
			t, err = api.PostTweet(tweet, nil)
			if err != nil {
				log.Error(err)
				continue
			}

			break
		}

		log.Infof("Success: %d.", t.Id)
		time.Sleep(time.Minute)
	}
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
	var ret string
	isFB := true
	if num%3 == 0 {
		ret = "Fizz"
	}

	if num%5 == 0 {
		ret += "Buzz"
	}

	if len(ret) == 0 {
		ret = fmt.Sprint(num)
		isFB = false
	}

	return ret, isFB
}
