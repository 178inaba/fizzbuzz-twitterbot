package mysql_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	"github.com/178inaba/fizzbuzz-twitterbot/model/mysql"
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"
)

type fizzbuzzTweetTestSuite struct {
	suite.Suite

	db      *sql.DB
	service mysql.FizzbuzzTweetService
}

func TestFizzbuzzTweetSuite(t *testing.T) {
	suite.Run(t, new(fizzbuzzTweetTestSuite))
}

func (s *fizzbuzzTweetTestSuite) SetupSuite() {
	db, err := mysql.Open("root", "fizzbuzz_twitterbot_test", true)
	s.NoError(err)

	s.db = db
	s.service = mysql.NewFizzbuzzTweetService(db)
}

func (s *fizzbuzzTweetTestSuite) SetupTest() {
	// Reset test db.
	_, err := s.db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.FizzbuzzTweetTableName))
	s.NoError(err)
	_, err = s.db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.NoError(err)
}

func (s *fizzbuzzTweetTestSuite) TestLatestTweet() {
	// No rows.
	tweet, err := s.service.LatestTweet()
	s.NoError(err)
	s.Nil(tweet)

	// Row exists.
	ft := &model.FizzbuzzTweet{Number: 1, Tweet: "test tweet!"}
	id, err := s.service.Insert(ft)
	s.NoError(err)

	tweet, err = s.service.LatestTweet()
	s.NoError(err)

	s.Equal(id, tweet.ID)
	s.Equal(ft.Number, tweet.Number)
	s.Equal(ft.Tweet, tweet.Tweet)
	s.Equal(uint64(0), tweet.TwitterTweetID)

	threeSecAgo := time.Now().UTC().Add(-3 * time.Second)
	s.True(tweet.UpdatedAt.After(threeSecAgo))
	s.True(tweet.CreatedAt.After(threeSecAgo))
}

func (s *fizzbuzzTweetTestSuite) TestInsert() {
	ft := &model.FizzbuzzTweet{Number: 1, Tweet: "test tweet!"}
	insertID, err := s.service.Insert(ft)
	s.NoError(err)
	s.Equal(uint64(1), insertID)

	rows, err := sq.Select("*").
		From(model.FizzbuzzTweetTableName).RunWith(s.db).Query()
	s.NoError(err)

	var cnt int
	for rows.Next() {
		var actual model.FizzbuzzTweet
		err := rows.Scan(&actual.ID, &actual.Number, &actual.Tweet,
			&actual.TwitterTweetID, &actual.UpdatedAt, &actual.CreatedAt)
		s.NoError(err)

		s.Equal(insertID, actual.ID)
		s.Equal(ft.Number, actual.Number)
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

func (s *fizzbuzzTweetTestSuite) TestAddTwitterTweetID() {
	// Error.
	err := s.service.AddTwitterTweetID(1, 1)
	s.Error(err)

	// No error.
	ft := &model.FizzbuzzTweet{Number: 1, Tweet: "test tweet!"}
	id, err := s.service.Insert(ft)
	s.NoError(err)

	err = s.service.AddTwitterTweetID(id, 1)
	s.NoError(err)
}
