package mysql_test

import (
	"testing"

	"github.com/178inaba/fizzbuzz-twitterbot/model"
	"github.com/178inaba/fizzbuzz-twitterbot/model/mysql"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db, err := mysql.Open("root:@/fizzbuzz_twitterbot_test?parseTime=true")
	assert.NoError(t, err)

	service := mysql.NewFizzbuzzTweetService(db)
	ft := &model.FizzbuzzTweet{Number: 1, Tweet: "test tweet!"}
	id, err := service.Insert(ft)
	assert.NoError(t, err)
	assert.True(t, id > 0)
}
