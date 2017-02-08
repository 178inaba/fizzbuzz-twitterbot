package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/178inaba/fizzbuzz-twitterbot/model/mysql"
	"github.com/178inaba/fizzbuzz-twitterbot/post"
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

func main() {
	kingpin.Parse()
	os.Exit(run())
}

func run() int {
	err := flagValidation()
	if err != nil {
		log.Errorf("Flag error: %s.", err)
		return 1
	}

	// Create twitter client.
	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	api := anaconda.NewTwitterApi(*accessToken, *accessTokenSecret)

	db, err := mysql.Open("root", "fizzbuzz_twitterbot", true)
	if err != nil {
		log.Errorf("DB open error: %s.", err)
		return 1
	}

	fts := mysql.NewFizzbuzzTweetService(db)
	pes := mysql.NewPostErrorService(db)
	c := post.NewClient(api, fts, pes)
	c.Post()

	return 0
}

func flagValidation() error {
	if *consumerKey == "" {
		return errors.New("consumer key is not set")
	}

	if *consumerSecret == "" {
		return errors.New("consumer secret is not set")
	}

	if *accessToken == "" {
		return errors.New("access token is not set")
	}

	if *accessTokenSecret == "" {
		return errors.New("access token secret is not set")
	}

	return nil
}
