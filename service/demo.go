package service

import (
	"errors"
	"github.com/1101137105/dcard/glob"
	"time"
)

type demoService struct{}

var DemoService = demoService{}
var (
	IpMap     = make(map[string]int)
	IpTimeMap = make(map[string]int64)
)

func (t demoService) CheckRequestLimit(ip string, limit int) (val int, err error) {

	var key string
	if _, ok := IpTimeMap[ip]; !ok {
		IpTimeMap[ip] = glob.GetNowUTC()
		key = glob.CreateIPKey(ip, glob.GetNowUTC())

	} else {
		subM := time.Unix(glob.GetNowUTC(), 0).Sub(time.Unix(IpTimeMap[ip], 0))
		if subM.Minutes() > 1 {
			key = glob.CreateIPKey(ip, glob.GetNowUTC())

		} else {
			key = glob.CreateIPKey(ip, IpTimeMap[ip])
		}
	}
	IpMap[key] = IpMap[key] + 1

	if IpMap[key] > limit {
		return 0, errors.New("Error")

	}

	return IpMap[key], nil

}
