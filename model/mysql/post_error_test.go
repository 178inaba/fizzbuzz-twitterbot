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

type postErrorSuite struct {
	suite.Suite

	db      *sql.DB
	service model.PostErrorService
}

func TestPostErrorSuite(t *testing.T) {
	suite.Run(t, new(postErrorSuite))
}

func (s *postErrorSuite) SetupSuite() {
	db, err := mysql.Open("root", "", "fizzbuzz_twitterbot_test")
	s.NoError(err)

	s.db = db
	s.service = mysql.NewPostErrorService(db)
}

func (s *postErrorSuite) SetupTest() {
	// Reset test db.
	_, err := s.db.Exec("SET foreign_key_checks = 0")
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.PostErrorTableName))
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.FizzbuzzTweetTableName))
	s.NoError(err)
	_, err = s.db.Exec("SET foreign_key_checks = 1")
	s.NoError(err)
}

func (s *postErrorSuite) TestInsert() {
	// Insert parent table.
	fts := mysql.NewFizzbuzzTweetService(s.db)
	ft := &model.FizzbuzzTweet{Number: math.MaxUint64, IsFizz: true, IsBuzz: true, Tweet: "FizzBuzz #18446744073709551615"}
	err := fts.Insert(ft)
	s.NoError(err)

	pe := &model.PostError{FizzbuzzTweetNumber: ft.Number, ErrorMessage: "test"}
	insertID, err := s.service.Insert(pe)
	s.NoError(err)
	s.Equal(uint64(1), insertID)

	rows, err := sq.Select("*").
		From(model.PostErrorTableName).RunWith(s.db).Query()
	s.NoError(err)

	var cnt int
	for rows.Next() {
		var actual model.PostError
		err := rows.Scan(&actual.ID, &actual.FizzbuzzTweetNumber,
			&actual.ErrorMessage, &actual.UpdatedAt, &actual.CreatedAt)
		s.NoError(err)

		s.Equal(insertID, actual.ID)
		s.Equal(ft.Number, actual.FizzbuzzTweetNumber)
		s.Equal(pe.ErrorMessage, actual.ErrorMessage)

		threeSecAgo := time.Now().UTC().Add(-3 * time.Second)
		s.True(actual.UpdatedAt.After(threeSecAgo))
		s.True(actual.CreatedAt.After(threeSecAgo))

		cnt++
	}

	s.Equal(1, cnt)
	s.NoError(rows.Err())
}

func (s *postErrorSuite) TearDownSuite() {
	s.db.Close()
}
