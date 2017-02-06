package mysql

import (
	"database/sql"
	"time"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	sq "github.com/Masterminds/squirrel"
)

// FizzbuzzTweetService is mysql database service.
type FizzbuzzTweetService struct {
	db *sql.DB
}

// NewFizzbuzzTweetService is create service.
func NewFizzbuzzTweetService(db *sql.DB) FizzbuzzTweetService {
	return FizzbuzzTweetService{db: db}
}

// NextNumber return next fizz buzz calculation number.
func (s FizzbuzzTweetService) NextNumber() (uint64, error) {
	var number, twitterTweetID uint64
	err := sq.Select("number", "twitter_tweet_id").From(model.FizzbuzzTweetTableName).
		OrderBy("updated_at desc").Limit(1).RunWith(s.db).Scan(&number, &twitterTweetID)
	if err != nil {
		return 0, err
	}

	if number == 0 {
		return 1, nil
	}

	if twitterTweetID == 0 {
		return number, nil
	}

	return number + 1, nil
}

// Insert is insert fizzbuzz_tweets table.
func (s FizzbuzzTweetService) Insert(f *model.FizzbuzzTweet) (uint64, error) {
	now := time.Now().UTC()
	res, err := sq.Insert(model.FizzbuzzTweetTableName).Columns(
		"number", "tweet", "updated_at", "created_at").
		Values(f.Number, f.Tweet, now, now).RunWith(s.db).Exec()
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
	_, err := sq.Update(model.FizzbuzzTweetTableName).
		Set("twitter_tweet_id", twitterTweetID).Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": id}).RunWith(s.db).Exec()
	if err != nil {
		return err
	}

	return nil
}
