package mysql_test

import (
	"testing"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	"github.com/178inaba/fizzbuzz-twitterbot/model/mysql"
	"github.com/stretchr/testify/suite"
)

type fizzbuzzTweetTestSuite struct {
	suite.Suite
	service mysql.FizzbuzzTweetService
}

func TestFizzbuzzTweetSuite(t *testing.T) {
	suite.Run(t, new(fizzbuzzTweetTestSuite))
}

func (s *fizzbuzzTweetTestSuite) SetupTest() {
	db, err := mysql.Open("root:@/fizzbuzz_twitterbot_test?parseTime=true")
	s.NoError(err)

	s.service = mysql.NewFizzbuzzTweetService(db)
}

func (s *fizzbuzzTweetTestSuite) TestInsert() {
	ft := &model.FizzbuzzTweet{Number: 1, Tweet: "test tweet!"}
	id, err := s.service.Insert(ft)
	s.NoError(err)
	s.True(id > 0)
}
