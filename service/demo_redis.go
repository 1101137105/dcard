package service

import (
	"github.com/1101137105/dcard/glob"
	"time"
)

type demoRedisService struct{}

var DemoRedisService = demoRedisService{}
var (
	IpRedisTimeMap = make(map[string]int64)
)

func (t demoRedisService) CheckRequestLimit(ip string, limit int) (val int, err error) {

	if _, ok := IpRedisTimeMap[ip]; !ok {
		IpRedisTimeMap[ip] = glob.GetNowUTC()
		if err = glob.RedisService.Set(ip, 1); err != nil {
			return
		}
		val = val + 1
		return
	}
	subM := time.Unix(glob.GetNowUTC(), 0).Sub(time.Unix(IpRedisTimeMap[ip], 0))
	if subM.Minutes() > 1 {
		IpRedisTimeMap[ip] = glob.GetNowUTC()

		if err = glob.RedisService.Set(ip, 0); err != nil {
			return
		}
	} else {
		if val, err = glob.RedisService.Get(ip); err != nil {

			return
		}
		if val >= limit {
			return
		}
		val = val + 1
		if err = glob.RedisService.Set(ip, val); err != nil {
			return
		}
	}

	return

}
