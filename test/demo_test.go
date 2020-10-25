package test

import (
	"errors"
	"github.com/1101137105/dcard/glob"
	"github.com/1101137105/dcard/service"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckRequestLimitOver(t *testing.T) {
	testIP := "1:1:1:1"
	var err error
	limit := 10
	var val int

	for i := 1; i < limit+1; i++ {
		val, err = service.DemoService.CheckRequestLimit(testIP, limit)
		if i > limit {
			assert.Equal(t, errors.New("error"), err)

		} else {
			assert.Equal(t, i, val)

		}
	}

	return
}
func TestCheckRequestRedisLimitOver(t *testing.T) {
	glob.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	testIP := "1:1:1:1"
	var err error
	limit := 60
	var val int
	for i := 1; i < limit+1; i++ {
		val, err = service.DemoRedisService.CheckRequestLimit(testIP, limit)
		if i > limit {
			assert.Equal(t, errors.New("error"), err)

		} else {
			assert.Equal(t, i, val)

		}
	}

	return

}
