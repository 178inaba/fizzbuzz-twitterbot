package model

import "time"

type fizzbuzzTweet struct {
	id             uint64
	number         uint64
	tweet          string
	twitterTweetID uint64
	updatedAt      time.Time
	createdAt      time.Time
}
