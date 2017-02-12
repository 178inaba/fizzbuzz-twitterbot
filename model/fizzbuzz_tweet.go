package model

import "time"

// FizzbuzzTweetTableName is fizzbuzz tweet table name.
const FizzbuzzTweetTableName = "fizzbuzz_tweets"

// FizzbuzzTweet is fizz buzz tweet object.
type FizzbuzzTweet struct {
	Number         uint64
	IsFizz         bool
	IsBuzz         bool
	Tweet          string
	TwitterTweetID uint64
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

// FizzbuzzTweetService is service interface.
type FizzbuzzTweetService interface {
	LatestTweet() (*FizzbuzzTweet, error)
	Insert(ft *FizzbuzzTweet) error
	AddTwitterTweetID(twitterTweetID, number uint64) error
}
