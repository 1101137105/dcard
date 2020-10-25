package glob

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

type redisService struct{}

var RedisService = redisService{}

func SetupRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Host + Config.Redis.Addr,
		Password: Config.Redis.Password, // no password set
		DB:       Config.Redis.Db,       // use default DB
	})

}

func (r *redisService) Set(key string, value int) (err error) {
	var ctx = context.Background()

	err = Rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return
	}
	return
}

func (r *redisService) Get(key string) (val int, err error) {
	var ctx = context.Background()

	val, err = Rdb.Get(ctx, key).Int()
	if err != nil {
		if err != redis.Nil {
			return

		} else {
			err = nil
		}
	}
	return
}
