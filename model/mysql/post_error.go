package mysql

import (
	"database/sql"
	"time"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	sq "github.com/Masterminds/squirrel"
)

// PostErrorService is mysql database service.
type PostErrorService struct {
	b sq.StatementBuilderType
}

// NewPostErrorService is create service.
func NewPostErrorService(db *sql.DB) PostErrorService {
	return PostErrorService{b: sq.StatementBuilder.RunWith(db)}
}

// Insert is insert post_errors table.
func (s PostErrorService) Insert(f *model.PostError) (uint64, error) {
	now := time.Now().UTC()
	res, err := s.b.Insert(model.FizzbuzzTweetTableName).Columns(
		"fizzbuzz_tweet_id", "error_message", "updated_at", "created_at").
		Values(f.FizzbuzzTweetID, f.ErrorMessage, now, now).Exec()
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
