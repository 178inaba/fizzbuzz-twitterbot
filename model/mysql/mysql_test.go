package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareExecer(t *testing.T) {
	db, err := Open("root", "", "fizzbuzz_twitterbot_test")
	assert.NoError(t, err)

	pe := prepareExecer{db: db}

	res, err := pe.Exec("prepare error")
	assert.Error(t, err)
	assert.Nil(t, res)
}
