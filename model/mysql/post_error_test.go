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

type postErrorTestSuite struct {
	suite.Suite

	db      *sql.DB
	service mysql.PostErrorService
}

func TestPostErrorSuite(t *testing.T) {
	suite.Run(t, new(postErrorTestSuite))
}

func (s *postErrorTestSuite) SetupSuite() {
	db, err := mysql.Open("root", "", "fizzbuzz_twitterbot_test")
	s.NoError(err)

	s.db = db
	s.service = mysql.NewPostErrorService(db)
}

func (s *postErrorTestSuite) SetupTest() {
	// Reset test db.
	_, err := s.db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.PostErrorTableName))
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.FizzbuzzTweetTableName))
	s.NoError(err)
	_, err = s.db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.NoError(err)
}

func (s *postErrorTestSuite) TestInsert() {
	// Insert parent table.
	fts := mysql.NewFizzbuzzTweetService(s.db)
	ft := &model.FizzbuzzTweet{Number: 1, Tweet: "test tweet!"}
	ftInsertID, err := fts.Insert(ft)
	s.NoError(err)

	pe := &model.PostError{FizzbuzzTweetID: ftInsertID, ErrorMessage: "Error!!"}
	insertID, err := s.service.Insert(pe)
	s.NoError(err)
	s.Equal(uint64(1), insertID)

	rows, err := sq.Select("*").
		From(model.PostErrorTableName).RunWith(s.db).Query()
	s.NoError(err)

	var cnt int
	for rows.Next() {
		var actual model.PostError
		err := rows.Scan(&actual.ID, &actual.FizzbuzzTweetID,
			&actual.ErrorMessage, &actual.UpdatedAt, &actual.CreatedAt)
		s.NoError(err)

		s.Equal(insertID, actual.ID)
		s.Equal(ftInsertID, actual.FizzbuzzTweetID)
		s.Equal(pe.ErrorMessage, actual.ErrorMessage)

		threeSecAgo := time.Now().UTC().Add(-3 * time.Second)
		s.True(actual.UpdatedAt.After(threeSecAgo))
		s.True(actual.CreatedAt.After(threeSecAgo))

		cnt++
	}

	s.Equal(1, cnt)
	s.NoError(rows.Err())
}

func (s *postErrorTestSuite) TearDownSuite() {
	s.db.Close()
}
