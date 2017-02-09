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

		ft := &model.FizzbuzzTweet{Number: i}
		err := c.post(ft)
		if err != nil {
			return err
		}
	}
}

func (c Client) post(ft *model.FizzbuzzTweet) error {
	ft.Tweet = tweetText(ft.Number)
	ftID, err := c.fts.Insert(ft)
	if err != nil {
		return err
	}

	c.logger.Printf("Tweet: %s.", ft.Tweet)
	var t anaconda.Tweet
	for {
		var err error
		t, err = c.api.PostTweet(ft.Tweet, nil)
		if err == nil {
			break
		}

		pe := &model.PostError{FizzbuzzTweetID: ftID, ErrorMessage: err.Error()}
		_, err = c.pes.Insert(pe)
		if err != nil {
			return err
		}

		c.logger.Printf("Error: %s.", err)
		time.Sleep(time.Second)
	}

	err = c.fts.AddTwitterTweetID(ftID, uint64(t.Id))
	if err != nil {
		return err
	}

	c.logger.Printf("Success! Twitter Tweet ID: %d.", t.Id)
	return nil
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
