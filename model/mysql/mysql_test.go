package mysql

import (
	"testing"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func TestPrepareExecer(t *testing.T) {
	db, err := Open("root", "", "fizzbuzz_twitterbot_test")
	assert.NoError(t, err)

	pe := prepareExecer{db: db}

	res, err := pe.Exec("prepare error")
	assert.Error(t, err)
	assert.Nil(t, res)

	sql, _, err := sq.Insert(model.FizzbuzzTweetTableName).
		Columns("number").Values(struct{}{}).ToSql()
	assert.NoError(t, err)

	res, err = pe.Exec(sql)
	assert.Error(t, err)
	assert.Nil(t, res)
}
