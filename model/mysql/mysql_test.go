package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"
)

type prepareExecerSuite struct {
	suite.Suite

	db     *sql.DB
	execer prepareExecer
}

func TestPrepareExecerSuite(t *testing.T) {
	suite.Run(t, new(prepareExecerSuite))
}

func (s *prepareExecerSuite) SetupSuite() {
	db, err := Open("root", "", "fizzbuzz_twitterbot_test")
	s.NoError(err)

	s.db = db
	s.execer = prepareExecer{db: db}
}

func (s *prepareExecerSuite) SetupTest() {
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

func (s *prepareExecerSuite) TestInsert() {
	res, err := s.execer.Exec("prepare error")
	s.Error(err)
	s.Nil(res)

	sql, _, err := sq.Insert(model.FizzbuzzTweetTableName).
		Columns("number").Values(struct{}{}).ToSql()
	s.NoError(err)

	res, err = s.execer.Exec(sql)
	s.Error(err)
	s.Nil(res)
}

func (s *prepareExecerSuite) TearDownSuite() {
	s.db.Close()
}
