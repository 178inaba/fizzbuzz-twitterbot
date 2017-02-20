package mysql

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

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

func (s *prepareExecerSuite) SetupTest() {
	db, err := Open("root", "", "fizzbuzz_twitterbot_test")
	s.NoError(err)

	s.db = db
	s.execer = prepareExecer{db: db}

	// Reset test db.
	_, err = s.db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.PostErrorTableName))
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", model.FizzbuzzTweetTableName))
	s.NoError(err)
	_, err = s.db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.NoError(err)
}

func (s *prepareExecerSuite) TestExec() {
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

func (s *prepareExecerSuite) TestOverMaxPrepare() {
	// Setup.
	showName := "max_prepared_stmt_count"
	row := s.db.QueryRow(fmt.Sprintf("SHOW GLOBAL VARIABLES LIKE '%s'", showName))

	var name string
	var cnt int
	err := row.Scan(&name, &cnt)
	s.NoError(err)
	s.Equal(showName, name)

	defer func() {
		// Tear down.
		_, err = s.db.Exec(fmt.Sprintf("SET GLOBAL max_prepared_stmt_count = %d", cnt))
		s.NoError(err)
	}()

	_, err = s.db.Exec("SET GLOBAL max_prepared_stmt_count = ?", 1)
	s.NoError(err)

	// Test.
	now := time.Now().UTC()
	ft := model.FizzbuzzTweet{}
	b := sq.Insert(model.FizzbuzzTweetTableName).Columns(
		"number", "is_fizz", "is_buzz", "tweet", "updated_at", "created_at")

	sql, args, err := b.Values(ft.Number, ft.IsFizz, ft.IsBuzz, ft.Tweet, now, now).ToSql()
	s.NoError(err)

	_, err = s.execer.Exec(sql, args...)
	s.NoError(err)

	sql, args, err = b.Values(ft.Number+1, ft.IsFizz, ft.IsBuzz, ft.Tweet, now, now).ToSql()
	s.NoError(err)

	_, err = s.execer.Exec(sql, args...)
	s.NoError(err)
}

func (s *prepareExecerSuite) TearDownTest() {
	s.db.Close()
}
