package model

import "time"

// PostErrorTableName is tweet post error log table name.
const PostErrorTableName = "post_errors"

// PostError is post error log.
type PostError struct {
	ID                  uint64
	FizzbuzzTweetNumber uint64
	ErrorMessage        string
	UpdatedAt           time.Time
	CreatedAt           time.Time
}

// PostErrorService is service interface.
type PostErrorService interface {
	Insert(pe *PostError) (uint64, error)
}
