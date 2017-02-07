package mysql_test

import (
	"database/sql"
	"fmt"
	"testing"

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
	db, err := mysql.Open("root:@/fizzbuzz_twitterbot_test?parseTime=true")
	s.NoError(err)

	s.db = db
	s.service = mysql.NewFizzbuzzTweetService(db)
}

func (s *fizzbuzzTweetTestSuite) SetupTest() {
	// Reset test db.
	_, err := sq.Delete(model.FizzbuzzTweetTableName).RunWith(s.db).Exec()
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", model.FizzbuzzTweetTableName))
	s.NoError(err)
}

func (s *fizzbuzzTweetTestSuite) TestInsert() {
	ft := &model.FizzbuzzTweet{Number: 1, Tweet: "test tweet!"}
	id, err := s.service.Insert(ft)
	s.NoError(err)
	s.Equal(uint64(1), id)
}

func (s *fizzbuzzTweetTestSuite) TestNextNumber() {
	number, err := s.service.NextNumber()
	s.NoError(err)
	s.Equal(uint64(1), number)
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
