package model

import "time"

type postError struct {
	id              uint64
	fizzbuzzTweetID uint64
	errorMessage    string
	updatedAt       time.Time
	createdAt       time.Time
}
