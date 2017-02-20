package mysql_test

import (
	"database/sql"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	"github.com/178inaba/fizzbuzz-twitterbot/model/mysql"
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"
)

type fizzbuzzTweetSuite struct {
	suite.Suite

	db      *sql.DB
	service model.FizzbuzzTweetService
}

func TestFizzbuzzTweetSuite(t *testing.T) {
	suite.Run(t, new(fizzbuzzTweetSuite))
}

func (s *fizzbuzzTweetSuite) SetupSuite() {
	db, err := mysql.Open("root", "", "fizzbuzz_twitterbot_test")
	s.NoError(err)

	s.db = db
	s.service = mysql.NewFizzbuzzTweetService(db)
}

func (s *fizzbuzzTweetSuite) SetupTest() {
	// Reset test db.
	_, err := s.db.Exec("SET foreign_key_checks = 0")
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.FizzbuzzTweetTableName))
	s.NoError(err)
	_, err = s.db.Exec("SET foreign_key_checks = 1")
	s.NoError(err)
}

func (s *fizzbuzzTweetSuite) TestLatestTweet() {
	// No rows.
	tweet, err := s.service.LatestTweet()
	s.NoError(err)
	s.Nil(tweet)

	// Row exists.
	ft := &model.FizzbuzzTweet{Number: math.MaxUint64, IsFizz: true, IsBuzz: true, Tweet: "FizzBuzz #18446744073709551615"}
	err = s.service.Insert(ft)
	s.NoError(err)

	tweet, err = s.service.LatestTweet()
	s.NoError(err)

	s.Equal(ft.Number, tweet.Number)
	s.Equal(ft.IsFizz, tweet.IsFizz)
	s.Equal(ft.IsBuzz, tweet.IsBuzz)
	s.Equal(ft.Tweet, tweet.Tweet)
	s.Equal(uint64(0), tweet.TwitterTweetID)

	threeSecAgo := time.Now().UTC().Add(-3 * time.Second)
	s.True(tweet.UpdatedAt.After(threeSecAgo))
	s.True(tweet.CreatedAt.After(threeSecAgo))
}

func (s *fizzbuzzTweetSuite) TestInsert() {
	ft := &model.FizzbuzzTweet{Number: math.MaxUint64, IsFizz: true, IsBuzz: true, Tweet: "Buzz #18446744073709551615"}
	err := s.service.Insert(ft)
	s.NoError(err)

	rows, err := sq.Select("*").
		From(model.FizzbuzzTweetTableName).RunWith(s.db).Query()
	s.NoError(err)

	var cnt int
	for rows.Next() {
		var actual model.FizzbuzzTweet
		err := rows.Scan(&actual.Number, &actual.IsFizz, &actual.IsBuzz,
			&actual.Tweet, &actual.TwitterTweetID, &actual.UpdatedAt, &actual.CreatedAt)
		s.NoError(err)

		s.Equal(ft.Number, actual.Number)
		s.Equal(ft.IsFizz, actual.IsFizz)
		s.Equal(ft.IsBuzz, actual.IsBuzz)
		s.Equal(ft.Tweet, actual.Tweet)
		s.Equal(uint64(0), actual.TwitterTweetID)

		threeSecAgo := time.Now().UTC().Add(-3 * time.Second)
		s.True(actual.UpdatedAt.After(threeSecAgo))
		s.True(actual.CreatedAt.After(threeSecAgo))

		cnt++
	}

	s.Equal(1, cnt)
	s.NoError(rows.Err())
}

func (s *fizzbuzzTweetSuite) TestAddTwitterTweetID() {
	// Error.
	err := s.service.AddTwitterTweetID(1, 1)
	s.Error(err)

	// No error.
	ft := &model.FizzbuzzTweet{Number: 15, IsFizz: true, IsBuzz: true, Tweet: "FizzBuzz #15"}
	err = s.service.Insert(ft)
	s.NoError(err)

	err = s.service.AddTwitterTweetID(math.MaxUint64, ft.Number)
	s.NoError(err)

	tweet, err := s.service.LatestTweet()
	s.NoError(err)
	s.Equal(uint64(math.MaxUint64), tweet.TwitterTweetID)
}

func (s *fizzbuzzTweetSuite) TearDownSuite() {
	s.db.Close()
}
