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
	var num uint64
	canInsert := true
	ft, err := c.fts.LatestTweet()
	if err != nil {
		return err
	} else if ft == nil {
		num = 1
	} else if ft.TwitterTweetID == 0 {
		num = ft.Number
		canInsert = false
	} else {
		num = ft.Number + 1
	}

	for i := uint64(num); ; i++ {
		// Next post to 00 second.
		waitNextZeroSec()

		err := c.post(i, canInsert)
		if err != nil {
			return err
		}

		canInsert = true
	}
}

func (c Client) post(num uint64, canInsert bool) error {
	isFizz, isBuzz := fizzbuzz(num)
	tweet := tweetText(isFizz, isBuzz, num)
	if canInsert {
		err := c.fts.Insert(&model.FizzbuzzTweet{Number: num, IsFizz: isFizz, IsBuzz: isBuzz, Tweet: tweet})
		if err != nil {
			return err
		}
	}

	c.logger.Printf("Tweet: %s.", tweet)
	var t anaconda.Tweet
	for {
		var err error
		t, err = c.api.PostTweet(tweet, nil)
		if err == nil {
			break
		}

		c.logger.Printf("Error: %s.", err)
		pe := &model.PostError{FizzbuzzTweetNumber: num, ErrorMessage: err.Error()}
		_, err = c.pes.Insert(pe)
		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	err := c.fts.AddTwitterTweetID(uint64(t.Id), num)
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

func tweetText(isFizz, isBuzz bool, num uint64) string {
	text := fizzbuzzText(isFizz, isBuzz)
	if len(text) > 0 {
		// Add number hashtag.
		return fmt.Sprintf("%s #%d", text, num)
	}

	return fmt.Sprint(num)
}

func fizzbuzzText(isFizz, isBuzz bool) string {
	var text string
	if isFizz {
		text = "Fizz"
	}

	if isBuzz {
		text += "Buzz"
	}

	return text
}

func fizzbuzz(num uint64) (isFizz, isBuzz bool) {
	if num%3 == 0 {
		isFizz = true
	}

	if num%5 == 0 {
		isBuzz = true
	}

	return isFizz, isBuzz
}
