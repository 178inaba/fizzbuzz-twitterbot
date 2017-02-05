package model

import "time"

// FizzbuzzTweetTableName is fizzbuzz tweet table name.
const FizzbuzzTweetTableName = "fizzbuzz_tweets"

// FizzbuzzTweet is fizz buzz tweet object.
type FizzbuzzTweet struct {
	ID             uint64
	Number         uint64
	Tweet          string
	TwitterTweetID uint64
	UpdatedAt      time.Time
	CreatedAt      time.Time
}
