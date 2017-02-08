package post

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
)

// Client is post client.
type Client struct {
	api    *anaconda.TwitterApi
	fts    model.FizzbuzzTweetService
	pes    model.PostErrorService
	logger logrus.StdLogger
}

// NewClient is create client struct.
func NewClient(api *anaconda.TwitterApi, fts model.FizzbuzzTweetService,
	pes model.PostErrorService, logger logrus.StdLogger) Client {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	return Client{api: api, fts: fts, pes: pes, logger: logger}
}

// Post is post fizz buzz tweet.
func (c Client) Post() error {
	var number uint64
	tweet, err := c.fts.LatestTweet()
	if err != nil {
		return err
	} else if tweet == nil {
		number = 1
	} else if tweet.TwitterTweetID == 0 {
		number = tweet.Number
	} else {
		number = tweet.Number + 1
	}

	for i := uint64(number); ; i++ {
		// Next post to 00 second.
		waitNextZeroSec()

		c.post(i)
	}
}

func (c Client) post(i uint64) {
	tweet := tweetText(i)
	c.logger.Printf("Tweet: %s.", tweet)
	var t anaconda.Tweet
	for {
		var err error
		t, err = c.api.PostTweet(tweet, nil)
		if err == nil {
			break
		}

		c.logger.Printf("Error: %s.", err)
		time.Sleep(time.Second)
	}

	c.logger.Printf("Success! Twitter Tweet ID: %d.", t.Id)
}

func nextZeroSec() time.Duration {
	n := time.Now()

	return n.Truncate(time.Minute).Add(time.Minute).Sub(n)
}

func waitNextZeroSec() {
	time.Sleep(nextZeroSec())
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
