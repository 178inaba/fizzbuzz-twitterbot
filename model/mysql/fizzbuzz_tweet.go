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
	b  sq.StatementBuilderType
	pe prepareExecer
}

// NewFizzbuzzTweetService is create service.
func NewFizzbuzzTweetService(db *sql.DB) FizzbuzzTweetService {
	return FizzbuzzTweetService{
		b: sq.StatementBuilder.RunWith(db), pe: prepareExecer{db: db}}
}

// LatestTweet return latest fizz buzz tweet.
func (s FizzbuzzTweetService) LatestTweet() (*model.FizzbuzzTweet, error) {
	ft := &model.FizzbuzzTweet{}
	err := s.b.Select("*").From(model.FizzbuzzTweetTableName).
		OrderBy("updated_at desc").Limit(1).
		Scan(&ft.Number, &ft.IsFizz, &ft.IsBuzz, &ft.Tweet,
			&ft.TwitterTweetID, &ft.UpdatedAt, &ft.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ft, nil
}

// Insert is insert fizzbuzz_tweets table.
func (s FizzbuzzTweetService) Insert(ft *model.FizzbuzzTweet) error {
	now := time.Now().UTC()
	sql, args, err := sq.Insert(model.FizzbuzzTweetTableName).Columns(
		"number", "is_fizz", "is_buzz", "tweet", "updated_at", "created_at").
		Values(ft.Number, ft.IsFizz, ft.IsBuzz, ft.Tweet, now, now).ToSql()
	if err != nil {
		return err
	}

	_, err = s.pe.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

// AddTwitterTweetID is update fizzbuzz_tweets table row of id.
func (s FizzbuzzTweetService) AddTwitterTweetID(twitterTweetID, number uint64) error {
	sql, args, err := sq.Update(model.FizzbuzzTweetTableName).
		Set("twitter_tweet_id", twitterTweetID).Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"number": number}).ToSql()
	if err != nil {
		return err
	}

	res, err := s.pe.Exec(sql, args...)
	if err != nil {
		return err
	}

	updateCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if updateCnt < 1 {
		return errors.Errorf("row not found: %d", number)
	}

	return nil
}
