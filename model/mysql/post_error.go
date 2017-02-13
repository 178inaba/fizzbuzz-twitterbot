package mysql

import (
	"database/sql"
	"time"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	sq "github.com/Masterminds/squirrel"
)

// PostErrorService is mysql database service.
type PostErrorService struct {
	pe prepareExecer
}

// NewPostErrorService is create service.
func NewPostErrorService(db *sql.DB) PostErrorService {
	return PostErrorService{pe: prepareExecer{db: db}}
}

// Insert is insert post_errors table.
func (s PostErrorService) Insert(pe *model.PostError) (uint64, error) {
	now := time.Now().UTC()
	sql, args, err := sq.Insert(model.PostErrorTableName).Columns(
		"fizzbuzz_tweet_number", "error_message", "updated_at", "created_at").
		Values(pe.FizzbuzzTweetNumber, pe.ErrorMessage, now, now).ToSql()
	if err != nil {
		return 0, err
	}

	res, err := s.pe.Exec(sql, args...)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
