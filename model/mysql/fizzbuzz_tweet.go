package mysql

import (
	"database/sql"
	"time"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// FizzbuzzTweetService is mysql database service.
type FizzbuzzTweetService struct {
	b sq.StatementBuilderType
}

// NewFizzbuzzTweetService is create service.
func NewFizzbuzzTweetService(db *sql.DB) FizzbuzzTweetService {
	return FizzbuzzTweetService{b: sq.StatementBuilder.RunWith(db)}
}

// NextNumber return next fizz buzz calculation number.
func (s FizzbuzzTweetService) NextNumber() (uint64, error) {
	var number, twitterTweetID uint64
	err := s.b.Select("number", "twitter_tweet_id").From(model.FizzbuzzTweetTableName).
		OrderBy("updated_at desc").Limit(1).Scan(&number, &twitterTweetID)
	if err == sql.ErrNoRows {
		return 1, nil
	} else if err != nil {
		return 0, err
	}

	if twitterTweetID == 0 {
		return number, nil
	}

	return number + 1, nil
}

// Insert is insert fizzbuzz_tweets table.
func (s FizzbuzzTweetService) Insert(ft *model.FizzbuzzTweet) (uint64, error) {
	now := time.Now().UTC()
	res, err := s.b.Insert(model.FizzbuzzTweetTableName).Columns(
		"number", "tweet", "updated_at", "created_at").
		Values(ft.Number, ft.Tweet, now, now).Exec()
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

// AddTwitterTweetID is update fizzbuzz_tweets table row of id.
func (s FizzbuzzTweetService) AddTwitterTweetID(id, twitterTweetID uint64) error {
	res, err := s.b.Update(model.FizzbuzzTweetTableName).
		Set("twitter_tweet_id", twitterTweetID).Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return err
	}

	updateCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if updateCnt < 1 {
		return errors.Errorf("row not found: %d", id)
	}

	return nil
}
