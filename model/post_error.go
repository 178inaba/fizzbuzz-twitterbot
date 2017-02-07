package model

import "time"

// PostError is post error log.
type PostError struct {
	ID              uint64
	FizzbuzzTweetID uint64
	ErrorMessage    string
	UpdatedAt       time.Time
	CreatedAt       time.Time
}
